package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"

	proto "github.com/ingirls/user/proto/user"
	"github.com/labstack/gommon/random"
	"golang.org/x/crypto/bcrypt"
)

// PasswordHash PasswordHash
func PasswordHash(password string) (string, string, error) {
	salt := random.String(6)
	md5Sign := Md5(fmt.Sprintf("%s%s", password, salt))
	bytes, err := bcrypt.GenerateFromPassword([]byte(md5Sign), bcrypt.DefaultCost)

	return string(bytes), salt, err
}

// PasswordVerify PasswordVerify
func PasswordVerify(password, salt, hash string) bool {
	md5Sign := Md5(fmt.Sprintf("%s%s", password, salt))
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(md5Sign))
	return err == nil
}

// Md5 Md5
func Md5(src string) string {
	h := md5.New()
	io.WriteString(h, src)
	return hex.EncodeToString(h.Sum(nil))
}

// ClearUserPassword ClearUserPassword
func ClearUserPassword(user *proto.User) *proto.User {
	user.Password = ""
	user.Salt = ""
	return user
}
