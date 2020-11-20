//Package context created by liuzhipeng04@meicai.cn @2019
package context
import (
	"github.com/gin-gonic/gin"
)

//HandlerFunc 为接受context扩展结构的方法
type HandlerFunc func(c *Context)

//Handle 方法为路由调用gin控制器方法
func Handle(h HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := &Context{
			c,
		}
		h(ctx)
	}
}

//Context 是对gin.Context进行的扩展结构
type Context struct {
	*gin.Context
}

//Success 是对Context扩展方法 接口请求成功时调用
func (c *Context) Success(code int, msg string, data interface{}) {
	c.JSON(200, gin.H{
		"ret":  1,
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

//Error 是对Context扩展方法 接口请求失败时调用
func (c *Context) Error(code int, msg string, data interface{}) {
	c.JSON(200, gin.H{
		"ret":  0,
		"code": code,
		"msg":  msg,
		"data": data,
	})
}