# Gin Context扩展包
## 1.包说明
本包扩展了`gin context`加入了以下两个函数,用于API请求时统一数据返回格式。
```go
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
```
## 2.使用方法
1.在基于gin项目的`router`包中引入本包
```go
import (
	ctxExt "git.sprucetec.com/golang-pkg/gin-context-ext"
	"github.com/gin-gonic/gin"
)
```  
可以根据需求使用`ctxExt`作为的`context`包的别名，防止包名冲突。

2.在路由方法中使用扩展包handle方法调用控制器方法。
```go
r.GET("/people/:id", ctxExt.Handle(peopleModule.GetPerson))
```  
3.在控制器中相应处理方法中接受`c *ctxExt.Context`扩展后的context为参数，然后调用扩展的方法。
```go
/*
获取人员
*/
func GetPerson(c *ctxExt.Context) {
	id := c.Params.ByName("id")
	fmt.Println(id)
	if person, err := peopleMod.GetPerson(id); err != nil {
		c.Error(1008, "无此记录！", "")
	} else {
		c.Success(200, "succ", person)
	}
}
```
