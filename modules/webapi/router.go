package webapi

import (
	"github.com/gin-gonic/gin"
)

// Route 路由信息
type Route struct {
	C Controller
}

// MountRoute 挂在路由信息
func (that *Route) MountRoute(r *gin.Engine) {
	g := r.Group("/webapi")
	// g.GET("/test/link-trace", that.C.TestLinkTrace)

	g.GET("/categorys", that.C.GetCategory)
	g.GET("/testok/:name/*action", func(c *gin.Context) {
		// c.String(http.StatusOK, "test is ok")
		// params 参数 /query
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		// query 参数 ?
		first := c.DefaultQuery("first", "guest") // 设置默认值
		second := c.Query("second")
		allName := "my name is " + first + " " + second
		c.JSON(200, gin.H{
			"code":    0,
			"message": message,
			"allName": allName,
		})
	})

}
