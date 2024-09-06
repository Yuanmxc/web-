package dao

import (
	"TTMS/kitex_gen/user"
	"context"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func getPassword(Password string) []byte {
	hash, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("bcrypt.GenerateFromPassword error")
		return nil
	}
	return hash
}
func PasswordEqual(Password1, Password2 string) error {
	//hash, _ := bcrypt.GenerateFromPassword([]byte(Password1), bcrypt.DefaultCost)
	err := bcrypt.CompareHashAndPassword([]byte(Password2), []byte(Password1))
	if err == nil { //密码正确
		return nil
	}
	return err
}

// CreateUser 创建用户最后要加上判断 userType=9的条件
func CreateUser(ctx context.Context, userInfo *user.User) (err error, id int64) {
	u := user.User{}
	if DB.Where("name = ?", userInfo.Name).Limit(1).Find(&u); u.Id > 0 {
		err = errors.New("[警告] 添加成功，您设置了重名用户")
	}
	u.Id = 0
	if DB.Where("email = ?", userInfo.Email).Limit(1).Find(&u); u.Id > 0 && len(u.Email) > 0 {
		err = errors.New("注册失败，该邮箱已经被使用")
		return
	}
	u.Id = 0
	userInfo.Password = string(getPassword(userInfo.Password))
	//fmt.Println("userInfo = ", userInfo, "\nu = ", u, "\nctx = ", ctx, "\nDB = ", DB)
	DB.Create(userInfo)
	DB.Select("id").Where("name = ?", userInfo.Name).Order("id desc ").Limit(1).Find(&u)
	id = u.Id
	return
}

func UserLogin(ctx context.Context, userInfo *user.User) (*user.User, error) {
	u := user.User{}
	DB.WithContext(ctx).Where("id = ?", userInfo.Id).Limit(1).Find(&u)
	if u.Id > 0 {
		if PasswordEqual(userInfo.Password, u.Password) == nil {
			return &u, nil
		} else {
			return &u, errors.New("密码错误")
		}
	}
	return &user.User{}, errors.New("未找到该用户")
}

func GetAllUser(ctx context.Context, current, pageSize int) ([]*user.User, int64, error) {
	users := make([]*user.User, pageSize)
	tx := DB.WithContext(ctx).Offset((current - 1) * pageSize).Limit(pageSize).Find(&users)
	for i, _ := range users {
		users[i].Password = "***"
	}
	var total int64
	tx.WithContext(ctx).Count(&total)
	return users, total, tx.Error
}
func ChangeUserPassword(ctx context.Context, UserId int, Password, NewPassword string) error {
	u := user.User{}
	DB.WithContext(ctx).Where("id = ?", UserId).Limit(1).Find(&u)
	if u.Id > 0 {
		if PasswordEqual(Password, u.Password) == nil {
			u.Password = string(getPassword(NewPassword))
			return DB.WithContext(ctx).Updates(&u).Error
		}
		return errors.New("原密码错误")
	} else {
		return errors.New("该用户不存在")
	}

}

func DeleteUser(ctx context.Context, UserId int) error {
	return DB.WithContext(ctx).Where("id = ?", UserId).Delete(&user.User{}).Error
}
func GetUserInfo(ctx context.Context, UserId int) (*user.User, error) {
	u := user.User{}
	tx := DB.WithContext(ctx).Where("id = ?", UserId).Limit(1).Find(&u)
	return &u, tx.Error
}

func BindEmail(ctx context.Context, id int64, email string) error {
	u := user.User{}
	DB.WithContext(ctx).Where("email = ?", email).Limit(1).Find(&u)
	if u.Id == 0 { //该邮箱尚未被其他用户绑定
		return DB.WithContext(ctx).Updates(&user.User{Id: id, Email: email}).Error
	} else if u.Id != id {
		return errors.New("该邮箱已经被其他用户绑定")
	}
	return nil
}
func ForgetPassword(ctx context.Context, Email string, NewPassword string) error {
	u := user.User{}
	DB.WithContext(ctx).Where("email = ?", Email).Limit(1).Find(&u)
	if u.Id > 0 {
		u.Password = string(getPassword(NewPassword))
		return DB.WithContext(ctx).Updates(&u).Error
	}
	return errors.New("该邮箱尚未绑定")
}
