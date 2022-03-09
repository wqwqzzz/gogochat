package ChatApi


import(
    "github.com/gin-gonic/gin"
    "gochat/internal/service"
    
    
    "net/http"
    "gochat/pkg/common/request"
    "gochat/pkg/common/response"
)

// 获取消息列表
func GetMessage(c *gin.Context) {
	var messageRequest request.MessageRequest
	err := c.BindQuery(&messageRequest)
	messages, err := service.MessageService.GetMessages(messageRequest)
	if err != nil {
		c.JSON(http.StatusOK, response.FailMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg(messages))
}
