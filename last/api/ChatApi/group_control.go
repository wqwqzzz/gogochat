package ChatApi


import(
    "github.com/gin-gonic/gin"
    "gochat/internal/service"
    "gochat/internal/model"
    
    "net/http"
    
    "gochat/pkg/common/response"
)
func GetGroup(c *gin.Context) {
	uuid := c.Param("uuid")
	groups, err := service.GroupService.GetGroups(uuid)
	if err != nil {
		c.JSON(http.StatusOK, response.FailMsg(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessMsg(groups))
}

// 保存分组列表
func SaveGroup(c *gin.Context) {
	uuid := c.Param("uuid")
	var group model.Group
	c.ShouldBindJSON(&group)

	service.GroupService.SaveGroup(uuid, group)
	c.JSON(http.StatusOK, response.SuccessMsg(nil))
}

// 加入组别
func JoinGroup(c *gin.Context) {
	userUuid := c.Param("userUuid")
	groupUuid := c.Param("groupUuid")
	err := service.GroupService.JoinGroup(groupUuid, userUuid)
	if err != nil {
		c.JSON(http.StatusOK, response.FailMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg(nil))
}

// 获取组内成员信息
func GetGroupUsers(c *gin.Context) {
	groupUuid := c.Param("uuid")
	users := service.GroupService.GetUserIdByGroupUuid(groupUuid)
	c.JSON(http.StatusOK, response.SuccessMsg(users))
}