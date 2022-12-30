package repos

import (
	"context"
	"errors"
	"fmt"
	"time"

	"dot.go/entities"
	"dot.go/helper"
	"gorm.io/gorm"
)

type RepoUserImpl struct {
	DB *gorm.DB
}

func NewRepoUserImpl(db *gorm.DB) RepoUser {
	return &RepoUserImpl{DB: db}
}

func (s RepoUserImpl) UserLogin(ctx context.Context, user entities.User) (int32, error) {
	funcNow := "RepoUserImpl.UserLogin"
	var u_email entities.User
	var u_pass entities.User

	email := user.Email
	pass := helper.HashPass(user.Password)

	tx := s.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		helper.LogError(err.Error(), "func:"+funcNow, "script: tx error")
		return 0, err
	}

	err := tx.WithContext(ctx).Find(&u_email, "email = ?", email)
	if err.Error != nil {
		tx.Rollback()
		helper.LogError(err.Error.Error(), "func:"+funcNow, "script: select user by email")
		return 0, err.Error
	}

	if u_email.ID == 0 {
		tx.Rollback()
		return 0, errors.New("Maaf, Email Atau Password Anda Salah")
	}

	err = tx.WithContext(ctx).Find(&u_pass, "password = ? and id = ?", pass, u_email.ID)
	if err.Error != nil {
		tx.Rollback()
		helper.LogError(err.Error.Error(), "func:"+funcNow, "script: select user by password")
		return 0, err.Error
	}

	if u_pass.ID == 0 {
		tx.Rollback()
		return 0, errors.New("Maaf, Email Atau Password Anda Salah")
	}

	formula := fmt.Sprintf("%v:%v:%v", time.Now().Unix(), helper.RandStringBytes(8), helper.RandStringNumber(8))
	hash := helper.EncSha1(formula)

	err = tx.WithContext(ctx).Model(&u_email).Update("device_id", hash)
	if err.Error != nil {
		tx.Rollback()
		helper.LogError(err.Error.Error(), "func:"+funcNow, "script: update user device id")
		return 0, err.Error
	}

	tx.Commit()

	return u_email.ID, nil
}

func (s RepoUserImpl) UserDetail(ctx context.Context, id_user int32) (entities.User, error) {
	funcNow := "RepoUserImpl.UserDetail"

	var u entities.User

	err := s.DB.WithContext(ctx).Find(&u, "id = ?", id_user)
	if err.Error != nil {
		helper.LogError(err.Error.Error(), "func:"+funcNow, "script: update user device id")
		return u, err.Error
	}

	return u, nil
}

func (s RepoUserImpl) UserGenerateToken(ctx context.Context, param ParamUser) (string, error) {
	funcNow := "RepoUserImpl.UserGenerateToken"

	uid := param.ID
	ua := param.UA
	data_user, err := s.UserDetail(ctx, uid)
	if err != nil {
		helper.LogError(err.Error(), "func:"+funcNow, "script: get detail user")
		return "", err
	}
	email := data_user.Email
	device := data_user.DeviceId
	formula := fmt.Sprintf("%v:%v:%v:%v", uid, email, device, ua)
	hash := helper.EncSha1(formula)
	return hash, nil
}

func (s RepoUserImpl) UserValidateToken(ctx context.Context, param ParamUser) (bool, error) {
	funcNow := "RepoUserImpl.UserValidateToken"

	id_user := param.ID
	token := param.Token
	ua := param.UA

	token_db, err := s.UserGenerateToken(ctx, ParamUser{ID: id_user, UA: ua})
	if err != nil {
		helper.LogError(err.Error(), "func:"+funcNow, "script: validasi token")
		return false, err
	}
	if token != token_db {
		return false, nil
	}
	return true, nil
}

func (s RepoUserImpl) UserUpdate(ctx context.Context, user entities.User) error {
	funcNow := "RepoUserImpl.UserUpdate"

	pass := helper.HashPass(user.Password)

	err := s.DB.WithContext(ctx).Table("users").Where("id = ?", user.ID).Updates(map[string]interface{}{"name": user.Name, "email": user.Email, "password": pass})
	if err.Error != nil {
		helper.LogError(err.Error.Error(), "func:"+funcNow, "script: update user")
		return err.Error
	}

	return nil
}
