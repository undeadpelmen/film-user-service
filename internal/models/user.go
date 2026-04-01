package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           uint
	Name         string
	PasswordHash string
	Role         string
}

func (u *User) String() string {
	if u == nil {
		return "user is nil"
	}

	str := "User: "

	if u.ID != 0 {
		str += fmt.Sprintf("ID: %d ", u.ID)
	}

	if u.Name != "" {
		str += fmt.Sprintf("Name: %s ", u.Name)
	}

	if u.PasswordHash != "" {
		str += fmt.Sprintf("PasswordHash: %s ", u.PasswordHash)
	}

	if u.Role != "" {
		str += fmt.Sprintf("Role: %s ", u.Role)
	}

	return str
}

func NewUser(name, passwordHash, role string) (uint, error) {
	user := &User{Name: name, PasswordHash: passwordHash, Role: role}

	result := dataBaseConnection.Create(user)

	if result.Error != nil {
		// TODO: rewite error handling
		return 0, result.Error
	}

	return user.ID, nil
}

func GetUserById(Id uint) (*User, error) {
	user := &User{ID: Id}

	result := dataBaseConnection.Find(user)

	if result.Error != nil {
		//TODO: rewrite error handling
		return nil, result.Error
	}

	return user, nil
}

func CheckUser(name, passwordHash string) (bool, *User, error) {
	user := &User{Name: name}

	result := dataBaseConnection.Where(user).Select("Name", "PasswordHash", "ID", "Role").Find(user)

	if result.Error != nil {
		//TODO: rewrite error handling
		return false, nil, result.Error
	}

	//log.Printf("cheking user with name: %s and password hash: %s, user: %v, user hash: %s", name, passwordHash, user, user.PasswordHash)

	if user.PasswordHash != passwordHash {
		return false, nil, nil
	}

	return true, user, nil
}

func DeleteUserById(Id uint) (*User, error) {
	user := &User{ID: Id}

	result := dataBaseConnection.Delete(user)

	if result.Error != nil {
		//TODO: rewrite error handling
		return nil, result.Error
	}

	return user, nil
}
