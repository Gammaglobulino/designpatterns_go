package flyweight_memory_save

import "strings"

type User struct {
	FullName string
}

//User factory
func NewUser(fullName string) *User {
	return &User{FullName: fullName}
}

var allNames []string

// this is the flyweight
type User2 struct {
	names []uint8
}

func NewUser2(fullName string) *User2 {
	getOrAdd := func(s string) uint8 {
		for i := range allNames {
			if allNames[i] == s {
				return uint8(i)
			}
		}
		allNames = append(allNames, s)
		return uint8(len(allNames) - 1)
	}

	result := User2{}
	parts := strings.Split(fullName, " ")
	for _, p := range parts {
		result.names = append(result.names, getOrAdd(p))
	}
	return &result
}

func (u *User2) FullName() string {
	var parts []string
	for _, id := range u.names {
		parts = append(parts, allNames[id])
	}
	return strings.Join(parts, " ")
}

func (u *User2) String() string {
	var parts []string
	for _, str := range u.names {
		parts = append(parts, string(str))
	}
	return strings.Join(parts, "")

}
