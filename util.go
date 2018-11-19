package util

import (
	"fmt"

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
func TokenEncrypt(key string, m map[string]interface{}) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)

	for index, val := range m {
		claims[index] = val
	}
	token.Claims = claims
	tokenString, err := token.SignedString([]byte(key))
	PanicOnError(err)
	return tokenString
}

// PasswordDecrypt 解析密钥
func PasswordDecrypt(tokenString string, key string) (interface{}, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		fmt.Println(err)
		return "", false
	}
}
