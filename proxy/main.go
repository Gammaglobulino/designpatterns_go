package proxy

import (
	"fmt"
)

type UserFinder interface {
	FindUser(id int32) (User, error)
}
type User struct {
	ID int32
}

type UserList []User

func (t *UserList) FindUser(id int32) (User, error) {

	for i := 0; i < len(*t); i++ {
		if (*t)[i].ID == id {
			return (*t)[i], nil
		}
	}
	return User{}, fmt.Errorf("User %d could not be found", id)
}

type UserListProxy struct {
	SomeDB               UserList
	StackCache           UserList
	StackCapacity        int
	LastSearchWasOnCache bool
}

func (up *UserListProxy) FindUser(id int32) (User, error) {
	user, err := up.StackCache.FindUser(id)
	if err == nil {
		fmt.Println("returning from the cache")
		up.LastSearchWasOnCache = true
		return user, nil
	}
	user, err = up.SomeDB.FindUser(id)
	if err != nil {
		return User{}, err
	}
	up.addUserToStack(user)
	fmt.Println("returning from the DB")
	up.LastSearchWasOnCache = false
	return user, nil
}

func (up *UserListProxy) addUserToStack(user User) {
	if len(up.StackCache) >= up.StackCapacity {
		up.StackCache = append(up.StackCache[1:], user)
	} else {
		up.StackCache.addUser(user)
	}
}

func (t *UserList) addUser(newUser User) {
	*t = append(*t, newUser)
}
