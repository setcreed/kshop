// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/setcreed/kshop/app/user/service/internal/data/ent/schema"
	"github.com/setcreed/kshop/app/user/service/internal/data/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[1].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescNickName is the schema descriptor for nick_name field.
	userDescNickName := userFields[2].Descriptor()
	// user.NickNameValidator is a validator for the "nick_name" field. It is called by the builders before save.
	user.NickNameValidator = userDescNickName.Validators[0].(func(string) error)
	// userDescHeadURL is the schema descriptor for head_url field.
	userDescHeadURL := userFields[3].Descriptor()
	// user.HeadURLValidator is a validator for the "head_url" field. It is called by the builders before save.
	user.HeadURLValidator = userDescHeadURL.Validators[0].(func(string) error)
	// userDescAddress is the schema descriptor for address field.
	userDescAddress := userFields[5].Descriptor()
	// user.AddressValidator is a validator for the "address" field. It is called by the builders before save.
	user.AddressValidator = userDescAddress.Validators[0].(func(string) error)
	// userDescRole is the schema descriptor for role field.
	userDescRole := userFields[8].Descriptor()
	// user.DefaultRole holds the default value on creation for the role field.
	user.DefaultRole = userDescRole.Default.(int)
}
