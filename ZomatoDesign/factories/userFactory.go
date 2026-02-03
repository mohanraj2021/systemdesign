package factories

import (
	"math/rand/v2"

	"github.com/zomatodesign/managers"
	"github.com/zomatodesign/models"
)

type UserFactory struct {
	users map[int]models.User
}

func (uf *UserFactory) CreateUser(user *models.User) models.User {
	user.Id = rand.Int()
	uf.users[user.Id] = *user
	return *user
}

func (uf *UserFactory) GetUser(id int) models.User {
	if user, exists := uf.users[id]; exists {
		return user
	}
	return models.User{}
}

func (uf *UserFactory) UpdateUser(id int, user *models.User) models.User {
	if _, exists := uf.users[id]; exists {
		uf.users[id] = *user
		return *user
	}
	return models.User{}
}

func (uf *UserFactory) DeleteUser(id int) bool {
	if _, exists := uf.users[id]; exists {
		delete(uf.users, id)
		return true
	}
	return false
}

func NewUserFactory() managers.UserManager {
	return &UserFactory{
		users: make(map[int]models.User),
	}
}
