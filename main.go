package main

import (
	"context"
	"fmt"
	"golang-started/config"
	_ "golang-started/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// _ "golang-started/docs"
	"golang-started/httperror"
	"golang-started/lib/opentracing"
	"golang-started/modules/bossapi"
	"golang-started/modules/webapi"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title teacher-example API
// @version v0.0.1
// @description 教师课件1.0
// @termOfService https://teacher-courseware
// @contract.name API.support
// @schemas http https
func main() {
	r := GetApp()
	listen := fmt.Sprintf(":%d", config.C.Port)
	srv := &http.Server{
		Addr:    listen,
		Handler: r,
	}
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("content-type", "application/json")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func GetApp() *gin.Engine {
	// 创建一个不包含中间件的路由器
	r := gin.New()
	// 自定义日志输出格式
	// TODO 解决输出Body
	r.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s %s\" headers:%s qs:%s\n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC1123),
			params.Method,
			params.Path,
			params.Request.Proto,
			params.StatusCode,
			params.Latency,
			params.ErrorMessage,
			params.Request.Header,
			params.Request.URL.RawQuery,
		)
	}))
	// TODO 需要解决自定义错误捕获
	r.Use(gin.Recovery())
	// 增加链路追踪
	r.Use(opentracing.GinOpentracingMiddleware())
	r.Use(httperror.Middleware())
	// cors解决跨域问题
	r.Use(Cors())
	// 装配webpai
	router := &webapi.Route{
		C: webapi.Controller{Service: &webapi.Service{}},
	}
	router.MountRoute(r)
	// 装配bossapi
	bossRouter := &bossapi.Route{
		C: bossapi.Controller{Service: &bossapi.Service{}},
	}
	bossRouter.MountRoute(r)
	urlStr := fmt.Sprintf("/swagger/doc.json")
	swaggerConfig := ginSwagger.URL(urlStr)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerConfig))
	// 控制台输出监听端口
	r.Run(":3301")
	// endless.ListenAndServe(":3301", r)
	return r
}
