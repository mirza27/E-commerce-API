package encrypt

import (
	"golang.org/x/crypto/bcrypt"
)



func MakePassword(pwd string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pwd), 14)
	if err != nil {
		panic(err)
	}

	result := string(hashed)
	return result, err
}


func CheckPassword(determinant []byte,pwd string) bool {
	pwd2 := []byte(pwd)
	isTrue := bcrypt.CompareHashAndPassword(determinant, pwd2)

	if isTrue == nil{
		return true
	} else {
		return false
	}
}