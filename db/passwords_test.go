package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/***** GetPasswordsForUser *****/
func TestGetPasswordsForUser(t *testing.T) {
	useTestSchema()

	passwords, err := GetPasswordsForUser(1)
	assert.Empty(t, err, "Should not get error with a valid id")

	expected := Password{
		Id:       1,
		Site:     "google.com",
		Login:    "denis",
		Password: "12345678",
	}
	assert.Equal(t, expected, passwords[0], "Should get the expected password")

	expected = Password{
		Id:       2,
		Site:     "facebook.com",
		Login:    "denis",
		Password: "87654321",
	}
	assert.Equal(t, expected, passwords[1], "Should get the expected password")
}

func TestGetPasswordsForNonExistingUser(t *testing.T) {
	useTestSchema()

	passwords, err := GetPasswordsForUser(0)
	assert.Empty(t, err, "Should not get error with a valid id")
	assert.Equal(t, 0, len(passwords), "Should not get any password for non-existing user")
}

/***** AddPassword *****/
func TestAddPasswordExistingUser(t *testing.T) {
	useTestSchema()

	id, err := AddPassword(1, "aaa.com", "aaa", "11111111")
	assert.Empty(t, err, "Should not get an error when adding a password to an existing user")
	assert.Greater(t, id, 0, "Should get a new password id greater than 0")

	nRows, err := DeletePassword(1, "aaa.com", "aaa")
	assert.Empty(t, err, "Should not get an error when removing an existing password")
	assert.Equal(t, 1, nRows, "Should have removed exactly 1 password")
}

func TestAddPasswordNonExistingUser(t *testing.T) {
	useTestSchema()

	_, err := AddPassword(0, "aaa.com", "aaa", "11111111")
	assert.NotEmpty(t, err, "Should get an error when adding a password to a non-existing user")
}

func TestAddPasswordExistingUserExistingPassword(t *testing.T) {
	useTestSchema()

	_, err := AddPassword(1, "google.com", "denis", "11111111")
	assert.NotEmpty(t, err, "Should get an error when adding an existing password to an existing user")
}

/***** DeletePassword *****/
func TestDeletePasswordNonExistingUser(t *testing.T) {
	useTestSchema()

	nRows, err := DeletePassword(0, "aaa.com", "aaa")
	assert.Empty(t, err, "Should not get an error when deleting a non-existing password")
	assert.Equal(t, 0, nRows, "Should not have deleted any password")
}

func TestDeletePasswordExistingUserNonExistingPassword(t *testing.T) {
	useTestSchema()

	nRows, err := UpdatePassword(1, "aaa.com", "aaa", "22222222")
	assert.Empty(t, err, "Should not get an error when deleting a non-existing password")
	assert.Equal(t, 0, nRows, "Should not have deleted any password")
}

/***** UpdatePassword *****/
func TestUpdatePasswordExistingUserNonExistingPassword(t *testing.T) {
	useTestSchema()

	nRows, err := UpdatePassword(1, "aaa.com", "aaa", "22222222")
	assert.Empty(t, err, "Should not get an error when updating a non-existing password")
	assert.Equal(t, 0, nRows, "Should not have updated any password")
}

func TestUpdatePasswordExistingUserExistingPassword(t *testing.T) {
	useTestSchema()

	id, err := AddPassword(1, "aaa.com", "aaa", "11111111")
	assert.Empty(t, err, "Should not get an error when adding a password to an existing user")
	assert.Greater(t, id, 0, "Should get a new password id greater than 0")

	nRows, err := UpdatePassword(1, "aaa.com", "aaa", "22222222")
	assert.Empty(t, err, "Should not get an error when updating an existing password")
	assert.Equal(t, 1, nRows, "Should have updated exactly 1 password")

	passwords, err := GetPasswordsForUser(1)
	assert.Empty(t, err, "Should not get error with a valid id")
	assert.Equal(t, 3, len(passwords), "Should get 3 passwords for user")
	assert.Equal(t, "22222222", passwords[2].Password, "Should get the updated password")

	nRows, err = DeletePassword(1, "aaa.com", "aaa")
	assert.Empty(t, err, "Should not get an error when removing an existing password")
	assert.Equal(t, 1, nRows, "Should have removed exactly 1 password")
}

func TestUpdatePasswordNonExistingUser(t *testing.T) {
	useTestSchema()

	nRows, err := UpdatePassword(0, "aaa.com", "aaa", "12121212")
	assert.Empty(t, err, "Should not get an error when updating a non-existing password")
	assert.Equal(t, 0, nRows, "Should not have updated any password")
}
