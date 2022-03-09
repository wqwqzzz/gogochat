package route
import(
    "gochat/api/ChatApi"
    "github.com/gin-gonic/gin"
    "net/http"
)


func NewRoute() *gin.Engine {
    gin.SetMode(gin.DebugMode)
    server:=gin.Default()
    server.Use(Cors())
    server.Use(gin.Recovery())
    group := server.Group("")
    {
        group.POST("/user/login", ChatApi.Login) //在前端中查看后，登录的路由就是这么写的，虽然这不符合由后端给前端吧，但是就这么看吧
        group.POST("/user/register", ChatApi.Register)
    }
    
    return server
    
}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				
			}
		}()

		c.Next()
	}
}