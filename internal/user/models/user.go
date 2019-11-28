package models

import (
	"github.com/payfazz/go-apt/pkg/fazzdb"
)

// User ...
type User struct {
	fazzdb.Model `json:"-"`
	Email        string `db:"email" json:"email"`
	RoleId       int64  `db:"role_id" json:"role_id"`
	IdentityCard string `db:"identity_card" json:"identity_card"`
	DivisionId   string `db:"division_id" json:"division_id"`
	FullName     string `db:"full_name" json:"full_name"`
}

// Payload is a function to read the payload data
func (m *User) Payload() map[string]interface{} {
	return m.MapPayload(m)
}

// Get ...
func (m *User) Get(key string) interface{} {
	return m.Payload()[key]
}

// UserModel is function to get User Model
func UserModel() *User {
	return &User{
		Model: fazzdb.PlainModel("users",
			[]fazzdb.Column{
				fazzdb.Col("email"),
				fazzdb.Col("role_id"),
				fazzdb.Col("identity_card"),
				fazzdb.Col("division_id"),
				fazzdb.Col("full_name"),
			},
			"email",
			true,
			true,
		),
	}
}