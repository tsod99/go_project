package api

// User
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Group    string `json:"group"`
}

// Group
type Group struct {
	Name string `json:"group_name"`
}

type UserField string

const (
	FieldUsername UserField = "username"
	FieldPassword           = "password"
	FieldEmail              = "email"
	FieldGroup              = "group"
)

// Check
func (uf UserField) Check() bool {
	return !(uf != FieldUsername && uf != FieldPassword && uf != FieldEmail && uf != FieldGroup)
}

func (uf UserField) String() string {
	return string(uf)
}

// UpdateUserRequest
type UpdateUserRequest struct {
	Username string    `json:"username"`
	Field    UserField `json:"field"`
	Value    string    `json:"value"`
}

// UpdateGroupRequest
type UpdateGroupRequest struct {
	GroupName    string `json:"group_name"`
	NewGroupName string `json:"new_group_name"`
}

// DeleteUserRequest
type DeleteUserRequest struct {
	Username string `json:"username"`
}

// DeleteGroupRequest
type DeleteGroupRequest struct {
	GroupName string `json:"group_name"`
}
