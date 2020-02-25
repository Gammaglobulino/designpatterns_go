package proxy

import (
	"../proxy"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestUserListProxy(t *testing.T) {
	someDatabase := proxy.UserList{}
	rand.Seed(2342342)
	for i := 0; i <= 1000000; i++ {
		n := rand.Int31()
		someDatabase = append(someDatabase, proxy.User{ID: n})
	}
	assert.EqualValues(t, 1000000, len(someDatabase)-1)

	proxy := proxy.UserListProxy{
		SomeDB:               someDatabase,
		StackCache:           proxy.UserList{},
		StackCapacity:        2,
		LastSearchWasOnCache: false,
	}
	knowwnIds := [3]int32{someDatabase[3].ID, someDatabase[4].ID, someDatabase[5].ID}
	t.Run("FindUser - Empty Cache", func(t *testing.T) {
		user, err := proxy.FindUser(knowwnIds[0])
		if err != nil {
			t.Fatal(err)
		}
		assert.EqualValues(t, user.ID, knowwnIds[0])
		assert.EqualValues(t, 1, len(proxy.StackCache))
		assert.False(t, proxy.LastSearchWasOnCache, "User should have been returned from an empty cache")
	})
	t.Run("FindUser - One user , ask for the same", func(t *testing.T) {
		user, err := proxy.FindUser(knowwnIds[0])
		if err != nil {
			t.Fatal(err)
		}
		assert.EqualValues(t, user.ID, knowwnIds[0])
		assert.True(t, proxy.LastSearchWasOnCache, "User should be returned from the cache")
	})
}
