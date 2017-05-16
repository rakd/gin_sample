package models


import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
)
// User ...
type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
	Verify   bool   `json:"verify"`
}


// Users ....
type Users []*User

// Len ...
func (s Users) Len() int {
	return len(s)
}

// NewUser ...
func NewUser() User{
	return User{
		Token:generageUUID(),
	}
}
// Create ...
func(u* User)Create()(User,error){
	origPassword := u.Password
	u.Password = hashedPassword(origPassword)
	err := db.Debug().Model(&u).Create(&u).Error
	if err!=nil{
		u.Password = origPassword
	}
	return *u,err
}




func hashedPassword(rawPassword string) string {
	s := sha256.New()
	s.Write([]byte(rawPassword))
	return base64.URLEncoding.EncodeToString(s.Sum(nil))
}





// Login ...
func (u *User) Login() (User, error) {
	var user User

	err := db.Where("email = ?", u.Email).Limit(1).First(&user).Error
	if err != nil {
		return user, errors.New("The username you entered doesn't belong to an account. Please check your username and try again")
	}

	err = db.Where("email = ?", u.Email).Where("password = ?", hashedPassword(u.Password)).Limit(1).First(&user).Error
	if err != nil {
		return user, errors.New("Sorry, your password was incorrect. Please double-check your password")
	}

	if user.ID == 0 {
		return user, errors.New("User not found")
	}

	return user, err
}