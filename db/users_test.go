package db

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
)

/***** GetUsers *****/
func TestGetUsers(t *testing.T) {
	useTestSchema()

	users, err := GetUsers()
	assert.Empty(t, err, "Should not get error when getting all users")
	// assuming at least 1 user
	assert.Equal(t, 2, len(users), "Should get exactly 2 users")

	expected := User{
		UserId:   1,
		UserName: "denis",
		FullName: sql.NullString{String: "Denis Rinfret", Valid: true},
		Email:    sql.NullString{String: "denis@profdenis.com", Valid: true},
	}
	assert.Equal(t, expected, users[0], "Should get the user expected")

	expected = User{
		UserId:   2,
		UserName: "alice",
		FullName: sql.NullString{String: "Alice Rinfret", Valid: true},
		Email:    sql.NullString{String: "", Valid: false},
	}
	assert.Equal(t, expected, users[1], "Should get the user expected")
}

/***** GetUserById *****/
func TestGetUserByIdExistingUser1(t *testing.T) {
	useTestSchema()

	user, err := GetUserById(1)
	assert.Empty(t, err, "Should not get error with a valid id")

	expected := User{
		UserId:   1,
		UserName: "denis",
		FullName: sql.NullString{String: "Denis Rinfret", Valid: true},
		Email:    sql.NullString{String: "denis@profdenis.com", Valid: true},
	}
	assert.Equal(t, expected, user, "Should get the user expected")
}

func TestGetUserByIdExistingUser2(t *testing.T) {
	useTestSchema()

	user, err := GetUserById(2)
	assert.Empty(t, err, "Should not get error with a valid id")

	expected := User{
		UserId:   2,
		UserName: "alice",
		FullName: sql.NullString{String: "Alice Rinfret", Valid: true},
		Email:    sql.NullString{String: "", Valid: false},
	}
	assert.Equal(t, expected, user, "Should get the user expected")
}

func TestGetUserByIdNonExistingUser(t *testing.T) {
	useTestSchema()

	_, err := GetUserById(0)
	assert.NotEmpty(t, err, "Should get an error with an invalid id")
}

/***** AddUser *****/
func TestAddUserExistingUser(t *testing.T) {
	useTestSchema()

	_, err := AddUser("denis", "Denis Rinfret", "12345678", "denis@profdenis.com")
	assert.NotEmpty(t, err, "Should get an error when adding an existing user")
}

func TestAddUserUserNameEmpty(t *testing.T) {
	useTestSchema()

	_, err := AddUser("", "Bob", "12345678", "")
	assert.NotEmpty(t, err, "Should get an error when adding a user with empty user name")
}

func TestAddUserNonExistingUser(t *testing.T) {
	useTestSchema()

	id, err := AddUser("bob", "Bob Bob", "12121212", "")
	assert.Empty(t, err, "Should not get an error when adding a new user")
	assert.Greater(t, id, 2, "Should get an id greater than 2 when adding a new user")

	// removing user to keep test reproducible
	nRows, err := DeleteUser(id)
	assert.Empty(t, err, "Should not get an error when removing an existing user")
	assert.Equal(t, 1, nRows, "Should have removed exactly 1 user")
}

/***** DeleteUser *****/
func TestDeleteUserNonExistingUser(t *testing.T) {
	useTestSchema()

	nRows, err := DeleteUser(0)
	assert.Empty(t, err, "Should not get an error when deleting a non-existing user")
	assert.Equal(t, 0, nRows, "Should not have deleted any user")
}

/***** UpdateUser *****/
func TestUpdateUserExistingUser(t *testing.T) {
	useTestSchema()

	// add 1 user, then update it, to keep test reproducible
	id, err := AddUser("bob", "Bob Bob", "12121212", "")
	assert.Empty(t, err, "Should not get an error when adding a new user")
	assert.Greater(t, id, 2, "Should get an id greater than 2 when adding a new user")

	nRows, err := UpdateUser(id, "bobbob", "Bob Bob", "12121212", "")
	assert.Empty(t, err, "Should not get an error when updating an existing user")
	assert.Equal(t, 1, nRows, "Should have updated exactly 1 user")

	// check that the update made it to the DB
	user, err := GetUserById(id)
	assert.Empty(t, err, "Should not get error with a valid id")
	assert.Equal(t, "bobbob", user.UserName, "Should get new username")

	// removing user to keep test reproducible
	nRows, err = DeleteUser(id)
	assert.Empty(t, err, "Should not get an error when removing an existing user")
	assert.Equal(t, 1, nRows, "Should have removed exactly 1 user")
}

func TestUpdateUserNonExistingUser(t *testing.T) {
	useTestSchema()

	nRows, err := UpdateUser(3, "bobbob", "Bob Bob", "12121212", "")
	assert.Empty(t, err, "Should not get an error when updating a non-existing user")
	assert.Equal(t, 0, nRows, "Should not have updated any user")
}
