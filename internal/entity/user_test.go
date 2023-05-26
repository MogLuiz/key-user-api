package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Luiz Henrique", "l@l.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "Luiz Henrique", user.Name)
	assert.Equal(t, "l@l.com", user.Email)
	assert.NotEmpty(t, user.Password)
}
