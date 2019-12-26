package users

import (
	"fmt"

	"github.com/waytkheming/bookstore-users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (u *User) Get() *errors.RestErr {
	result := usersDB[u.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", u.ID))
	}
	u.ID = result.ID
	u.FirstName = result.FirstName
	u.LastName = result.LastName
	u.Email = result.Email
	u.DateCreated = result.DateCreated
	return nil
}
func (u *User) Save() *errors.RestErr {
	newUser := usersDB[u.ID]
	if newUser != nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d already exists", u.ID))
	}
	newUser = u
	return nil
}
