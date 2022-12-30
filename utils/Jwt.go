package utils

import (
	"errors"
	"fmt"
	"time"

	"dot.go/config"
	"dot.go/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func GetIDUser(c *fiber.Ctx) (string, error) {
	data, err := GetDataFromJwt(c)
	if err != nil {
		return "", err
	}
	id_user := fmt.Sprintf("%v", data["id"])
	return id_user, nil
}

type FullDataJwt struct {
	Id    string
	Token string
}

func GetFullDataJwt(c *fiber.Ctx) (FullDataJwt, error) {
	var res FullDataJwt
	data, err := GetDataFromJwt(c)
	if err != nil {
		return res, err
	}
	id_member := fmt.Sprintf("%v", data["id"])
	token := fmt.Sprintf("%v", data["token"])
	res.Id = id_member
	res.Token = token
	return res, nil
}

func GetDataFromJwt(c *fiber.Ctx) (jwt.MapClaims, error) {
	tokenBytes := c.Request().Header.Peek("Authorization")
	tokenHeader := string(tokenBytes[:])
	if tokenHeader == "" {
		return nil, errors.New("Token Kosong")
	}
	tokenJwt := tokenHeader
	errValidate := ValidateJwt(tokenJwt)
	if errValidate != nil {
		return nil, errValidate
	}

	conf := config.GetMyConfig()
	jwtSecret := conf.APP.JWT_SECRET
	hmacSampleSecret := []byte(jwtSecret)
	tokenString := tokenJwt
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func ValidateJwt(tokenJwt string) error {
	conf := config.GetMyConfig()
	jwtSecret := conf.APP.JWT_SECRET
	var tokenString = tokenJwt

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if token.Valid {
		return nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return err
		} else {
			return err
		}
	} else {
		return err
	}
}

func RefreshTokenJwt(id int, level int, token string) (string, string, error) {
	conf := config.GetMyConfig()
	jwtSecret := conf.APP.JWT_SECRET
	mySigningKey := []byte(jwtSecret)

	tokenJWt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      id,
		"token":   token,
		"expired": helper.IntToString(int(time.Now().Add(time.Hour * (24 * 5)).Unix())),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenLogin, err := tokenJWt.SignedString(mySigningKey)
	if err != nil {
		return "", "", err
	}

	tokenJWtRefresh := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      id,
		"token":   token,
		"expired": helper.IntToString(int(time.Now().Add(time.Hour * (24 * 8)).Unix())),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenLoginRefresh, err := tokenJWtRefresh.SignedString(mySigningKey)
	if err != nil {
		return "", "", err
	}

	return tokenLogin, tokenLoginRefresh, nil
}

type ParamGenerateJwtTokens struct {
	ID    int
	Token string
}

func GenerateJwtToken(param ParamGenerateJwtTokens) (string, string, error) {
	id := param.ID
	token := param.Token
	conf := config.GetMyConfig()
	jwtSecret := conf.APP.JWT_SECRET
	mySigningKey := []byte(jwtSecret)

	tokenJWt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      id,
		"token":   token,
		"expired": helper.IntToString(int(time.Now().Add(time.Hour * (24 * 5)).Unix())),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenLogin, err := tokenJWt.SignedString(mySigningKey)
	if err != nil {
		return "", "", err
	}

	tokenJWtRefresh := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      id,
		"token":   token,
		"expired": helper.IntToString(int(time.Now().Add(time.Hour * (24 * 8)).Unix())),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenLoginRefresh, err := tokenJWtRefresh.SignedString(mySigningKey)
	if err != nil {
		return "", "", err
	}

	return tokenLogin, tokenLoginRefresh, nil
}
