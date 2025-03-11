//JWT工具类 （生成token解析token 获取当前登录用户id及其用户信息）

package jwt

import (
	"admin-go-api/api/entity"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type userStdClaims struct {
	entity.JwtAdmin
	jwt.RegisteredClaims // StandardClaims更替
}

// 过期时间

const TokenExpireDuration = time.Hour * 24

//token 密钥

var Secret = []byte("admin-go-api")
var (
	ErrAbsent  = "token absent"  // 令牌不存在
	ErrInvalid = "token invalid" // 令牌无效
)

//跟模用户信息生成token

func GenerateTokenByAdmin(admin entity.SysAdmin) (string, error) {
	var jwtAdmin = entity.JwtAdmin{
		ID:       admin.ID,
		Username: admin.Username,
		Nickname: admin.Nickname,
		Icon:     admin.Icon,
		Email:    admin.Email,
		Note:     admin.Note,
	}
	c := userStdClaims{
		jwtAdmin, // 自定义字段
		jwt.RegisteredClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			// 过期时间
			Issuer: "admin", //签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(Secret)

}

// ValidateToken 解析JWT
func ValidateToken(tokenString string) (*entity.JwtAdmin, error) {
	if tokenString == "" {
		return nil, errors.New(ErrAbsent)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if token == nil {
		return nil, errors.New(ErrInvalid)
	}
	claims := userStdClaims{}
	_, err = jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	return &claims.JwtAdmin, err
}
