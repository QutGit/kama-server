package opentracing

import (
	"context"
	"fmt"
	"golang-started/config"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func init() {
	env := os.Getenv("GO_ENV")
	os.Setenv("JAEGER_SAMPLER_TYPE", "ratelimiting")
	os.Setenv("JAEGER_SAMPLER_PARAM", "5")
	switch env {
	case "stage":
		os.Setenv("JAEGER_ENDPOINT", "http://tracing-analysis-dc-bj-internal.aliyuncs.com/adapt_fpfh1gftus@8bb412d1cd53424_fpfh1gftus@53df7ad2afe8301/api/traces")
	case "production":
		os.Setenv("JAEGER_ENDPOINT", "http://tracing-analysis-dc-hz-internal.aliyuncs.com/adapt_fpfh1gftus@8bb412d1cd53424_fpfh1gftus@53df7ad2afe8301/api/traces")
	default:
		os.Setenv("JAEGER_ENDPOINT", "http://10.8.8.210:14268/api/traces")
		os.Setenv("JAEGER_SAMPLER_TYPE", "const")
		os.Setenv("JAEGER_SAMPLER_PARAM", "1")
	}
	cfg, err := jaegercfg.FromEnv()
	if len(cfg.ServiceName) == 0 {
		cfg.ServiceName = fmt.Sprintf("%s-%s", config.C.Name, env)
	}

	if err != nil {
		log.Printf("Could not parse Jaeger env vars: %s", err.Error())
		return
	}
	tracer, _, err := cfg.NewTracer()
	if err != nil {
		log.Printf("Could not initialize jaeger tracer: %s", err.Error())
		return
	}
	opentracing.SetGlobalTracer(tracer)
}

func GinOpentracingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.URL.Path
		ctx, err := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(c.Request.Header))
		var span opentracing.Span
		if err != nil {
			span = opentracing.StartSpan(url)
		} else {
			span = opentracing.StartSpan(url, ext.RPCServerOption(ctx))
		}
		ext.HTTPUrl.Set(span, c.Request.URL.String())
		ext.HTTPMethod.Set(span, c.Request.Method)
		// 写入tag
		cNew := opentracing.ContextWithSpan(c.Request.Context(), span)
		c.Request = c.Request.WithContext(cNew)
		// 将新的链路信息注入到 c.Request.Header 中
		defer span.Finish()
		c.Next()
		status := uint16(c.Writer.Status())
		ext.HTTPStatusCode.Set(span, status)
		if status >= 400 {
			// 报错时，将tag标记为错误，同时记录调试信息
			ext.Error.Set(span, true)
			span.LogKV("body", c.Request.Body)
		}
	}
}

// DoWithLinkTrace
//
// 兼容 http.Client Do function
// 增加啦链路追踪
//
// <code>
// 	c := &lib.HttpClient{}
//	httpReq, _ := http.NewRequest("GET", "http://localhost:3002/health", nil)
//	c.DoWithLinkTrace(ctxNew, httpReq)
// </code>
//
func DoWithLinkTrace(ctx context.Context, r *http.Request) (context.Context, *http.Response, error) {
	c := &http.Client{}
	span, ctxNew := opentracing.StartSpanFromContext(ctx, r.URL.String())
	defer span.Finish()
	ext.SpanKindRPCClient.Set(span)
	ext.HTTPMethod.Set(span, r.Method)
	ext.HTTPUrl.Set(span, r.URL.String())
	opentracing.GlobalTracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(r.Header))
	resp, err := c.Do(r)
	if err != nil {
		span.SetTag(string(ext.Error), true)
	} else {
		if resp.StatusCode >= 400 {
			span.SetTag(string(ext.Error), true)
		}
		span.SetTag(string(ext.HTTPStatusCode), resp.StatusCode)
	}
	return ctxNew, resp, err
}
