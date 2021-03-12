package bossapi

import (
	"github.com/gin-gonic/gin"
)

// Route 路由信息
type Route struct {
	C Controller
}

// MountRoute 挂在路由信息
func (that *Route) MountRoute(r *gin.Engine) {
	g := r.Group("/bossapi")
	g.GET("/categorys", that.C.GetCategory)
	g.POST("/create", that.C.CreateArticle)
	g.GET("/articles", that.C.GetArticles)
}
