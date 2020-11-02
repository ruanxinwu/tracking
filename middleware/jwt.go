package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"tracking/global"
	"tracking/global/response"
	"tracking/model"
	"tracking/model/request"
	"tracking/service"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")

		modelToken := model.JwtBlackList{Jwt: token}

		if token == "" {
			response.Result(response.ERROR, gin.H{
				"reload": true,
			}, "未登陆或者非法登陆", c)
			c.Abort()
			return
		}

		if service.IsBlackList(token, modelToken) {
			response.Result(response.ERROR, gin.H{
				"reload": true,
			}, "您的账户异地登陆或者令牌失效", c)
			c.Abort()
			return
		}

		j := NewJwt()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				response.Result(response.ERROR, gin.H{
					"reload": true,
				}, "授权已过期", c)
				c.Abort()
				return
			}
			response.Result(response.ERROR, gin.H{
				"reload": true,
			}, err.Error(), c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}

type Jwt struct {
	SigningKey []byte
}

var (
	TokenMalformed   = errors.New("That's not even a token")
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJwt() *Jwt {
	return &Jwt{
		[]byte(global.GVA_CONFIG.Jwt.SigningKey),
	}
}

func (j *Jwt) ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, request.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	}
	return nil, TokenInvalid
}
