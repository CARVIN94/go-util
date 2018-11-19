package util

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo/bson"
	"golang.org/x/crypto/bcrypt"
)

// PanicOnError 当错误时输出并停止运行
func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

// ObjectIds []string to []bson.ObjectId
func ObjectIds(strArr []string) (arr []bson.ObjectId) {
	if len(strArr) == 0 {
		return arr
	}
	for _, v := range strArr {
		arr = append(arr, bson.ObjectIdHex(v))
	}
	return arr
}

// PasswordEncrypt 密码加密
func PasswordEncrypt(pwd string) []byte {
	password, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	PanicOnError(err)
	return password
}

// PasswordCompare 密码对比
func PasswordCompare(pwd string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(pwd))
	result := true
	if err != nil {
		result = false
	}
	return result
}

// TokenEncrypt 生成密钥
func TokenEncrypt(m jwt.MapClaims, key string) (string, bool) {
	t := time.Now()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = t.Add(time.Hour * 24 * 3).Unix()
	claims["iat"] = t.Unix()
	for index, val := range m {
		claims[index] = val
	}
	token.Claims = claims
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", false
	}
	return tokenString, true
}

// TokenDecrypt 解析密钥
func TokenDecrypt(tokenString string, key string) (jwt.MapClaims, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	}
	fmt.Println(err)
	return jwt.MapClaims{}, false
}
