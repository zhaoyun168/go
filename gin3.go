package main
 
import "github.com/gin-gonic/gin"
 
func main() {
    //Default返回一个默认的路由引擎
    r := gin.Default()
    r.GET("/user/info", func(c *gin.Context) {
        //输出json结果给调用方
        c.JSON(200, gin.H{
            "message": "pong get",
        })
    })
 
    r.POST("/user/info", func(c *gin.Context) {
        //输出json结果给调用方
        c.JSON(200, gin.H{
            "message": "pong post",
        })
    })
 
    r.PUT("/user/info", func(c *gin.Context) {
        //输出json结果给调用方
        c.JSON(200, gin.H{
            "message": "pong put",
        })
    })
 
    r.DELETE("/user/info", func(c *gin.Context) {
        //输出json结果给调用方
        c.JSON(200, gin.H{
            "message": "pong delete",
        })
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}