package pwd

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// 哈希密码
func HashPwd(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
 
// 验证密码  哈希之后的密码   输入的密码
func CheckPwd(hashPwd string, pwd string) bool {
	byteHash := []byte(hashPwd)
 
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(pwd))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
