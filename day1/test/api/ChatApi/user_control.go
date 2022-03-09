package ChatApi
import(
    "github.com/gin-gonic/gin"
    "gochat/internal/service"
    "gochat/internal/model"
    
    "net/http"
   
    "gochat/pkg/common/response"
)
func Login(c *gin.Context){
    var user model.User
    //然后绑定到json上
    c.ShouldBindJSON(&user)
    
    if service.UserService.Login(&user){
        c.JSON(http.StatusOK,response.SuccessMsg(user))
        return
    }
    c.JSON(http.StatusOK, response.FailMsg("Login failed"))
    
    
}
func Register(c *gin.Context){
    var user model.User
    c.ShouldBindJSON(&user)
    
    err := service.UserService.Register(&user)
    if err != nil {
    	c.JSON(http.StatusOK, response.FailMsg(err.Error()))
    	return
    }
    
    c.JSON(http.StatusOK, response.SuccessMsg(user))
    
}