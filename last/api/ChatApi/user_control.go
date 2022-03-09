package ChatApi
import(
    "github.com/gin-gonic/gin"
    "gochat/internal/service"
    "gochat/internal/model"
    
    "net/http"
   "gochat/pkg/common/request"
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

func AddFriend(c *gin.Context){
    var userFriendRequest request.FriendRequest
    c.ShouldBindJSON(&userFriendRequest)
    err := service.UserService.AddFriend(&userFriendRequest)
    if err != nil{
        c.JSON(http.StatusOK,response.FailMsg(err.Error()))
        return 
    }
    c.JSON(http.StatusOK,response.SuccessMsg(nil))
}
func ModifyUserInfo(c *gin.Context) {
	var user model.User
	c.ShouldBindJSON(&user)
	if err := service.UserService.ModifyUserInfo(&user); err != nil {
		c.JSON(http.StatusOK, response.FailMsg(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessMsg(nil))
}

func GetUserDetails(c *gin.Context) {
	uuid := c.Param("uuid")

	c.JSON(http.StatusOK, response.SuccessMsg(service.UserService.GetUserDetails(uuid)))
}

// 通过用户名获取用户信息
func GetUserOrGroupByName(c *gin.Context) {
	name := c.Query("name")

	c.JSON(http.StatusOK, response.SuccessMsg(service.UserService.GetUserOrGroupByName(name)))
}

func GetUserList(c *gin.Context) {
	uuid := c.Query("uuid")
	c.JSON(http.StatusOK, response.SuccessMsg(service.UserService.GetUserList(uuid)))
}