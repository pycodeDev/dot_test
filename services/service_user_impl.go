package services

import (
	"context"
	"time"

	"dot.go/config"
	"dot.go/entities"
	"dot.go/helper"
	"dot.go/model"
	"dot.go/rcode"
	"dot.go/repos"
	"dot.go/utils"
	"github.com/gofiber/fiber/v2"
)

type ServiceUserImpl struct{}

func NewServiceUserImpl() ServiceUser {
	return &ServiceUserImpl{}
}

func (s ServiceUserImpl) UserLogin(c *fiber.Ctx, user entities.User, ua string) error {
	ctx := context.Background()
	db := config.GormConnectWrite()

	rd := config.Redislocal()
	defer rd.Close()

	key_redis := rcode.PASS_WRONG_MEMBER_5_TIMES + user.Email + ":" + helper.GetCurrentDateFull()
	count_wrong_pass, _ := rd.Get(ctx, key_redis).Result()
	if helper.StringToInt(count_wrong_pass) > 3 {
		return model.ReturnError(c, "Maaf Anda telah Salah Input Password Sebanyak 3x", rcode.INTERNAL_ERROR)
	}

	id_user, err := repos.NewRepoUserImpl(db).UserLogin(ctx, user)
	if err != nil {
		count := helper.StringToInt(count_wrong_pass) + 1
		rd.Set(ctx, key_redis, count, time.Minute*3)
		return model.ReturnError(c, err.Error(), rcode.UNAUTHORIZED)
	}

	detail_user, err := repos.NewRepoUserImpl(db).UserDetail(ctx, id_user)
	if err != nil {
		return model.ReturnError(c, err.Error(), rcode.INTERNAL_ERROR)
	}

	token_user, err := repos.NewRepoUserImpl(db).UserGenerateToken(ctx, repos.ParamUser{ID: id_user, UA: ua})
	if err != nil {
		return model.ReturnError(c, err.Error(), rcode.INTERNAL_ERROR)
	}

	token_jwt, refresh_token, err := utils.GenerateJwtToken(utils.ParamGenerateJwtTokens{
		ID:    int(id_user),
		Token: token_user,
	})

	if err != nil {
		return model.ReturnError(c, err.Error(), rcode.INTERNAL_ERROR)
	}

	detail_user.Password = "******"
	detail_user.DeviceId = "------"

	jwt := fiber.Map{"token": token_jwt, "token_refresh": refresh_token}
	data := fiber.Map{"jwt": jwt, "user": detail_user}

	return c.JSON(model.SuccessResponse{
		Status:    1,
		Rc:        rcode.RESPONSE_SUCCESS,
		Message:   "Success Login !!",
		Data:      data,
		TimeStamp: time.Now().Unix(),
	})
}

func (s ServiceUserImpl) UserValidateToken(id_user int32, token string, ua string) (bool, error) {
	ctx := context.Background()
	db := config.GormConnectWrite()
	dbsql := config.MySQLConnect()
	defer dbsql.Close()

	repo_user := repos.NewRepoUserImpl(db)
	return repo_user.UserValidateToken(ctx, repos.ParamUser{
		ID:    id_user,
		Token: token,
		UA:    ua,
	})
}

func (s ServiceUserImpl) UserUpdate(c *fiber.Ctx, user entities.User) error {
	ctx := context.Background()
	db := config.GormConnectWrite()

	err := repos.NewRepoUserImpl(db).UserUpdate(ctx, user)
	if err != nil {
		return model.ReturnError(c, err.Error(), rcode.INTERNAL_ERROR)
	}

	return model.ReturnDefault(c, "Data Success Update", rcode.RESPONSE_SUCCESS)
}
