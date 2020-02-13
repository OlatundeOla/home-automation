// Code generated by jrpc. DO NOT EDIT.

package userdef

import (
	time "time"
)

// User is defined in the .def file
type User struct {
	Id        uint32    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Validate returns an error if any of the fields have bad values
func (m *User) Validate() error {

	return nil
}

// GetUserRequest is defined in the .def file
type GetUserRequest struct {
	UserId uint32 `json:"user_id"`
}

// Validate returns an error if any of the fields have bad values
func (m *GetUserRequest) Validate() error {

	return nil
}

// GetUserResponse is defined in the .def file
type GetUserResponse struct {
	User User `json:"user"`
}

// Validate returns an error if any of the fields have bad values
func (m *GetUserResponse) Validate() error {
	if err := m.User.Validate(); err != nil {
		return err
	}

	return nil
}

// ListUsersRequest is defined in the .def file
type ListUsersRequest struct {
}

// Validate returns an error if any of the fields have bad values
func (m *ListUsersRequest) Validate() error {
	return nil
}

// ListUsersResponse is defined in the .def file
type ListUsersResponse struct {
	Users []User `json:"users"`
}

// Validate returns an error if any of the fields have bad values
func (m *ListUsersResponse) Validate() error {
	for _, r := range m.Users {
		if err := r.Validate(); err != nil {
			return err
		}
	}

	return nil
}
