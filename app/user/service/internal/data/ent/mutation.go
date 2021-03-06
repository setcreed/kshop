// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/setcreed/kshop/app/user/service/internal/data/ent/predicate"
	"github.com/setcreed/kshop/app/user/service/internal/data/ent/user"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeUser = "User"
)

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	mobile        *string
	password_hash *string
	nick_name     *string
	head_url      *string
	birthday      *time.Time
	address       *string
	desc          *string
	gender        *int
	addgender     *int
	role          *int
	addrole       *int
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetMobile sets the "mobile" field.
func (m *UserMutation) SetMobile(s string) {
	m.mobile = &s
}

// Mobile returns the value of the "mobile" field in the mutation.
func (m *UserMutation) Mobile() (r string, exists bool) {
	v := m.mobile
	if v == nil {
		return
	}
	return *v, true
}

// OldMobile returns the old "mobile" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldMobile(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldMobile is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldMobile requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMobile: %w", err)
	}
	return oldValue.Mobile, nil
}

// ResetMobile resets all changes to the "mobile" field.
func (m *UserMutation) ResetMobile() {
	m.mobile = nil
}

// SetPasswordHash sets the "password_hash" field.
func (m *UserMutation) SetPasswordHash(s string) {
	m.password_hash = &s
}

// PasswordHash returns the value of the "password_hash" field in the mutation.
func (m *UserMutation) PasswordHash() (r string, exists bool) {
	v := m.password_hash
	if v == nil {
		return
	}
	return *v, true
}

// OldPasswordHash returns the old "password_hash" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldPasswordHash(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldPasswordHash is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldPasswordHash requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPasswordHash: %w", err)
	}
	return oldValue.PasswordHash, nil
}

// ResetPasswordHash resets all changes to the "password_hash" field.
func (m *UserMutation) ResetPasswordHash() {
	m.password_hash = nil
}

// SetNickName sets the "nick_name" field.
func (m *UserMutation) SetNickName(s string) {
	m.nick_name = &s
}

// NickName returns the value of the "nick_name" field in the mutation.
func (m *UserMutation) NickName() (r string, exists bool) {
	v := m.nick_name
	if v == nil {
		return
	}
	return *v, true
}

// OldNickName returns the old "nick_name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldNickName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldNickName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldNickName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldNickName: %w", err)
	}
	return oldValue.NickName, nil
}

// ClearNickName clears the value of the "nick_name" field.
func (m *UserMutation) ClearNickName() {
	m.nick_name = nil
	m.clearedFields[user.FieldNickName] = struct{}{}
}

// NickNameCleared returns if the "nick_name" field was cleared in this mutation.
func (m *UserMutation) NickNameCleared() bool {
	_, ok := m.clearedFields[user.FieldNickName]
	return ok
}

// ResetNickName resets all changes to the "nick_name" field.
func (m *UserMutation) ResetNickName() {
	m.nick_name = nil
	delete(m.clearedFields, user.FieldNickName)
}

// SetHeadURL sets the "head_url" field.
func (m *UserMutation) SetHeadURL(s string) {
	m.head_url = &s
}

// HeadURL returns the value of the "head_url" field in the mutation.
func (m *UserMutation) HeadURL() (r string, exists bool) {
	v := m.head_url
	if v == nil {
		return
	}
	return *v, true
}

// OldHeadURL returns the old "head_url" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldHeadURL(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldHeadURL is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldHeadURL requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldHeadURL: %w", err)
	}
	return oldValue.HeadURL, nil
}

// ClearHeadURL clears the value of the "head_url" field.
func (m *UserMutation) ClearHeadURL() {
	m.head_url = nil
	m.clearedFields[user.FieldHeadURL] = struct{}{}
}

// HeadURLCleared returns if the "head_url" field was cleared in this mutation.
func (m *UserMutation) HeadURLCleared() bool {
	_, ok := m.clearedFields[user.FieldHeadURL]
	return ok
}

// ResetHeadURL resets all changes to the "head_url" field.
func (m *UserMutation) ResetHeadURL() {
	m.head_url = nil
	delete(m.clearedFields, user.FieldHeadURL)
}

// SetBirthday sets the "birthday" field.
func (m *UserMutation) SetBirthday(t time.Time) {
	m.birthday = &t
}

// Birthday returns the value of the "birthday" field in the mutation.
func (m *UserMutation) Birthday() (r time.Time, exists bool) {
	v := m.birthday
	if v == nil {
		return
	}
	return *v, true
}

// OldBirthday returns the old "birthday" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldBirthday(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldBirthday is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldBirthday requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBirthday: %w", err)
	}
	return oldValue.Birthday, nil
}

// ClearBirthday clears the value of the "birthday" field.
func (m *UserMutation) ClearBirthday() {
	m.birthday = nil
	m.clearedFields[user.FieldBirthday] = struct{}{}
}

// BirthdayCleared returns if the "birthday" field was cleared in this mutation.
func (m *UserMutation) BirthdayCleared() bool {
	_, ok := m.clearedFields[user.FieldBirthday]
	return ok
}

// ResetBirthday resets all changes to the "birthday" field.
func (m *UserMutation) ResetBirthday() {
	m.birthday = nil
	delete(m.clearedFields, user.FieldBirthday)
}

// SetAddress sets the "address" field.
func (m *UserMutation) SetAddress(s string) {
	m.address = &s
}

// Address returns the value of the "address" field in the mutation.
func (m *UserMutation) Address() (r string, exists bool) {
	v := m.address
	if v == nil {
		return
	}
	return *v, true
}

// OldAddress returns the old "address" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldAddress(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldAddress is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldAddress requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAddress: %w", err)
	}
	return oldValue.Address, nil
}

// ClearAddress clears the value of the "address" field.
func (m *UserMutation) ClearAddress() {
	m.address = nil
	m.clearedFields[user.FieldAddress] = struct{}{}
}

// AddressCleared returns if the "address" field was cleared in this mutation.
func (m *UserMutation) AddressCleared() bool {
	_, ok := m.clearedFields[user.FieldAddress]
	return ok
}

// ResetAddress resets all changes to the "address" field.
func (m *UserMutation) ResetAddress() {
	m.address = nil
	delete(m.clearedFields, user.FieldAddress)
}

// SetDesc sets the "desc" field.
func (m *UserMutation) SetDesc(s string) {
	m.desc = &s
}

// Desc returns the value of the "desc" field in the mutation.
func (m *UserMutation) Desc() (r string, exists bool) {
	v := m.desc
	if v == nil {
		return
	}
	return *v, true
}

// OldDesc returns the old "desc" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldDesc(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldDesc is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldDesc requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDesc: %w", err)
	}
	return oldValue.Desc, nil
}

// ClearDesc clears the value of the "desc" field.
func (m *UserMutation) ClearDesc() {
	m.desc = nil
	m.clearedFields[user.FieldDesc] = struct{}{}
}

// DescCleared returns if the "desc" field was cleared in this mutation.
func (m *UserMutation) DescCleared() bool {
	_, ok := m.clearedFields[user.FieldDesc]
	return ok
}

// ResetDesc resets all changes to the "desc" field.
func (m *UserMutation) ResetDesc() {
	m.desc = nil
	delete(m.clearedFields, user.FieldDesc)
}

// SetGender sets the "gender" field.
func (m *UserMutation) SetGender(i int) {
	m.gender = &i
	m.addgender = nil
}

// Gender returns the value of the "gender" field in the mutation.
func (m *UserMutation) Gender() (r int, exists bool) {
	v := m.gender
	if v == nil {
		return
	}
	return *v, true
}

// OldGender returns the old "gender" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldGender(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldGender is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldGender requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldGender: %w", err)
	}
	return oldValue.Gender, nil
}

// AddGender adds i to the "gender" field.
func (m *UserMutation) AddGender(i int) {
	if m.addgender != nil {
		*m.addgender += i
	} else {
		m.addgender = &i
	}
}

// AddedGender returns the value that was added to the "gender" field in this mutation.
func (m *UserMutation) AddedGender() (r int, exists bool) {
	v := m.addgender
	if v == nil {
		return
	}
	return *v, true
}

// ClearGender clears the value of the "gender" field.
func (m *UserMutation) ClearGender() {
	m.gender = nil
	m.addgender = nil
	m.clearedFields[user.FieldGender] = struct{}{}
}

// GenderCleared returns if the "gender" field was cleared in this mutation.
func (m *UserMutation) GenderCleared() bool {
	_, ok := m.clearedFields[user.FieldGender]
	return ok
}

// ResetGender resets all changes to the "gender" field.
func (m *UserMutation) ResetGender() {
	m.gender = nil
	m.addgender = nil
	delete(m.clearedFields, user.FieldGender)
}

// SetRole sets the "role" field.
func (m *UserMutation) SetRole(i int) {
	m.role = &i
	m.addrole = nil
}

// Role returns the value of the "role" field in the mutation.
func (m *UserMutation) Role() (r int, exists bool) {
	v := m.role
	if v == nil {
		return
	}
	return *v, true
}

// OldRole returns the old "role" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldRole(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldRole is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldRole requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRole: %w", err)
	}
	return oldValue.Role, nil
}

// AddRole adds i to the "role" field.
func (m *UserMutation) AddRole(i int) {
	if m.addrole != nil {
		*m.addrole += i
	} else {
		m.addrole = &i
	}
}

// AddedRole returns the value that was added to the "role" field in this mutation.
func (m *UserMutation) AddedRole() (r int, exists bool) {
	v := m.addrole
	if v == nil {
		return
	}
	return *v, true
}

// ResetRole resets all changes to the "role" field.
func (m *UserMutation) ResetRole() {
	m.role = nil
	m.addrole = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *UserMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *UserMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *UserMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *UserMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *UserMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *UserMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 11)
	if m.mobile != nil {
		fields = append(fields, user.FieldMobile)
	}
	if m.password_hash != nil {
		fields = append(fields, user.FieldPasswordHash)
	}
	if m.nick_name != nil {
		fields = append(fields, user.FieldNickName)
	}
	if m.head_url != nil {
		fields = append(fields, user.FieldHeadURL)
	}
	if m.birthday != nil {
		fields = append(fields, user.FieldBirthday)
	}
	if m.address != nil {
		fields = append(fields, user.FieldAddress)
	}
	if m.desc != nil {
		fields = append(fields, user.FieldDesc)
	}
	if m.gender != nil {
		fields = append(fields, user.FieldGender)
	}
	if m.role != nil {
		fields = append(fields, user.FieldRole)
	}
	if m.created_at != nil {
		fields = append(fields, user.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, user.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldMobile:
		return m.Mobile()
	case user.FieldPasswordHash:
		return m.PasswordHash()
	case user.FieldNickName:
		return m.NickName()
	case user.FieldHeadURL:
		return m.HeadURL()
	case user.FieldBirthday:
		return m.Birthday()
	case user.FieldAddress:
		return m.Address()
	case user.FieldDesc:
		return m.Desc()
	case user.FieldGender:
		return m.Gender()
	case user.FieldRole:
		return m.Role()
	case user.FieldCreatedAt:
		return m.CreatedAt()
	case user.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldMobile:
		return m.OldMobile(ctx)
	case user.FieldPasswordHash:
		return m.OldPasswordHash(ctx)
	case user.FieldNickName:
		return m.OldNickName(ctx)
	case user.FieldHeadURL:
		return m.OldHeadURL(ctx)
	case user.FieldBirthday:
		return m.OldBirthday(ctx)
	case user.FieldAddress:
		return m.OldAddress(ctx)
	case user.FieldDesc:
		return m.OldDesc(ctx)
	case user.FieldGender:
		return m.OldGender(ctx)
	case user.FieldRole:
		return m.OldRole(ctx)
	case user.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case user.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldMobile:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMobile(v)
		return nil
	case user.FieldPasswordHash:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPasswordHash(v)
		return nil
	case user.FieldNickName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetNickName(v)
		return nil
	case user.FieldHeadURL:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetHeadURL(v)
		return nil
	case user.FieldBirthday:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBirthday(v)
		return nil
	case user.FieldAddress:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAddress(v)
		return nil
	case user.FieldDesc:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDesc(v)
		return nil
	case user.FieldGender:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetGender(v)
		return nil
	case user.FieldRole:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRole(v)
		return nil
	case user.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case user.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	var fields []string
	if m.addgender != nil {
		fields = append(fields, user.FieldGender)
	}
	if m.addrole != nil {
		fields = append(fields, user.FieldRole)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case user.FieldGender:
		return m.AddedGender()
	case user.FieldRole:
		return m.AddedRole()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	case user.FieldGender:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddGender(v)
		return nil
	case user.FieldRole:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddRole(v)
		return nil
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(user.FieldNickName) {
		fields = append(fields, user.FieldNickName)
	}
	if m.FieldCleared(user.FieldHeadURL) {
		fields = append(fields, user.FieldHeadURL)
	}
	if m.FieldCleared(user.FieldBirthday) {
		fields = append(fields, user.FieldBirthday)
	}
	if m.FieldCleared(user.FieldAddress) {
		fields = append(fields, user.FieldAddress)
	}
	if m.FieldCleared(user.FieldDesc) {
		fields = append(fields, user.FieldDesc)
	}
	if m.FieldCleared(user.FieldGender) {
		fields = append(fields, user.FieldGender)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	switch name {
	case user.FieldNickName:
		m.ClearNickName()
		return nil
	case user.FieldHeadURL:
		m.ClearHeadURL()
		return nil
	case user.FieldBirthday:
		m.ClearBirthday()
		return nil
	case user.FieldAddress:
		m.ClearAddress()
		return nil
	case user.FieldDesc:
		m.ClearDesc()
		return nil
	case user.FieldGender:
		m.ClearGender()
		return nil
	}
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldMobile:
		m.ResetMobile()
		return nil
	case user.FieldPasswordHash:
		m.ResetPasswordHash()
		return nil
	case user.FieldNickName:
		m.ResetNickName()
		return nil
	case user.FieldHeadURL:
		m.ResetHeadURL()
		return nil
	case user.FieldBirthday:
		m.ResetBirthday()
		return nil
	case user.FieldAddress:
		m.ResetAddress()
		return nil
	case user.FieldDesc:
		m.ResetDesc()
		return nil
	case user.FieldGender:
		m.ResetGender()
		return nil
	case user.FieldRole:
		m.ResetRole()
		return nil
	case user.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case user.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown User edge %s", name)
}
