package httperror

import (
	"fmt"

	logger "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Middleware 处理 HTTPError 错误
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			err := recover()
			if err == nil {
				return
			}
			logger.Error(err)
			// 如果是自定义错误，捕获并且抛出自定义错误
			if IsHTTPError(err) {
				v, _ := err.(*HTTPError)
				c.JSON(v.StatusCode, err)
				return
			}
			v, ok := err.(validator.ValidationErrors)
			if ok {
				e := BadRequest(v.Error(), 4000)
				c.JSON(e.StatusCode, e)
				return
			}
			message := fmt.Sprintf("%v", err)
			e := InternalError(message, 9999)
			c.JSON(e.StatusCode, e)
		}()
		c.Next()
	}
}
