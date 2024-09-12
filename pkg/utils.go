package pkg

import "golang.org/x/crypto/bcrypt"

const Salt = "gin-demo"

func HashPassword(password string) string {
	saltedPassword := password + Salt
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)
	return string(hashedPassword)
}

func CheckPassword(hashedPassword, password string) error {
	saltedPassword := password + Salt
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(saltedPassword))
}
