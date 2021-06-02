package userDao

import (
	"gorm.io/gorm"
	"kNotice/app/common/model"
	"kNotice/utils/db"
)

type userDao struct {
	Db *gorm.DB
}

func NewUserDao() *userDao {
	return &userDao{
		Db: db.NewDB(),
	}
}

//GetUserById 通过id获取用户
func (t *userDao) GetUserById(id int) (*model.UserModel, error) {
	user := &model.UserModel{}
	Db := t.Db
	res := Db.First(user, id)
	return user, res.Error
}

//CreateUser 创建用户
func (t *userDao) CreateUser(user *model.UserModel) (user2 *model.UserModel, err error) {
	Db := t.Db
	res := Db.Create(&user)
	return user, res.Error
}

//GetAllUser 获取全部用户
func (t *userDao) GetAllUser() (users []*model.UserModel, err error) {
	Db := t.Db
	users = make([]*model.UserModel, 0)
	res := Db.Where("status = ?", 1).Find(&users)
	return users, res.Error
}
