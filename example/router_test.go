package example

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func getHttpHandler() http.Handler {
	r := gin.New()
	route := &Route{
		C: Controller{
			Service: &Service{},
		},
	}
	route.MountRoute(r)
	return r
}

func TestExampleLinkTrace(t *testing.T) {
	app := getHttpHandler()
	req := httptest.NewRequest("GET", "/example/test/link-trace", nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)
	resp := w.Result()
	fmt.Println(resp.StatusCode)
	assert.Equal(t, resp.StatusCode, 200)
}

func TestExampleCreate(t *testing.T) {
	app := getHttpHandler()
	req := httptest.NewRequest("POST", "/example/", strings.NewReader("{\"name\":\"test\"}"))
	req.Header.Set("UID", "1")
	req.Header.Set("content-type", "application/json")
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)
	resp := w.Result()
	assert.Equal(t, resp.StatusCode, 200)
}
