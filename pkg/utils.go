package pkg

import "crypto/md5"

const Salt = "gin-demo"

func EncryptPassword(pwd string) string {
	hash := md5.New()
	d := pwd + Salt
	hash.Write([]byte(d))
	sum := hash.Sum(nil)
	return string(sum)
}
