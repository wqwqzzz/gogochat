package service
import (
    "gochat/internal/model"
    "gochat/internal/pool"
    "gochat/pkg/errors"
    "github.com/google/uuid"
    "time"
    
)
type userService struct {
}
var UserService = new(userService)

func (u *userService) Login(user *model.User) bool {
	pool.GetDB().AutoMigrate(&user)
	db := pool.GetDB()
	var queryUser *model.User
	db.First(&queryUser, "username = ?", user.Username)
	user.Uuid = queryUser.Uuid
    
	return queryUser.Password == user.Password
}
func (u *userService) Register(user *model.User) error {
    
    db := pool.GetDB()
    var userCount int64 //用usercount来判断id是否重复
    db.Model(user).Where("username", user.Username).Count(&userCount)
    if userCount>0{
        return errors.New("user exists!")
    }
    user.Uuid = uuid.New().String()//随机生成uuid
    user.CreateAt = time.Now()
    user.DeleteAt = 0
    
    db.Create(&user)
    return nil
    
    
}