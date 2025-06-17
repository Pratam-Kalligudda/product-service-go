package helper

import (
	"errors"
	"reflect"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

type Helper struct {
	Secret string
}

func (h Helper) Authorize(ctx fiber.Ctx) error {
	refreshToken := ctx.Cookies("refresh-token", "")
	if refreshToken == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "no refresh token",
		})
	}

	_, err := h.VerifyToken(refreshToken)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid refresh token :" + err.Error(),
		})
	}

	tokenArr := ctx.GetReqHeaders()["Authorization"]
	if len(tokenArr) < 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid token",
		})
	}
	tokenSep := strings.Split(tokenArr[0], " ")
	if tokenSep[0] != "Bearer" || len(tokenSep[1]) <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid token",
		})
	}

	token := tokenSep[1]
	claims, err := h.VerifyToken(token)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid token : " + err.Error(),
		})
	}

	userId, ok := claims["sub"].(float64)
	if !ok {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid token no subject in token :",
		})
	}

	ctx.Locals("userId", uint(userId))
	return ctx.Next()
}

func (h Helper) VerifyToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid siging method" + reflect.TypeOf(*token).Name())
		}
		return []byte(h.Secret), nil
	})
	if err != nil {
		return nil, errors.New("error while verfiying token + " + h.Secret + " :" + err.Error())
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}
	return nil, errors.New("error while finding claims")
}
