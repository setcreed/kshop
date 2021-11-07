package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("mobile").Unique().Comment("手机号码"),
		field.String("password_hash").Comment("密码"),
		field.String("nick_name").MaxLen(20).Comment("昵称").Optional(),
		field.String("head_url").MaxLen(200).Comment("头像").Optional(),
		field.Time("birthday").Comment("生日").SchemaType(map[string]string{
			dialect.MySQL: "date",
		}).Optional(),
		field.String("address").MaxLen(200).Comment("地址").Optional(),
		field.Text("desc").Comment("描述信息").Optional(),
		field.Int("gender").Comment("性别").Optional(),
		field.Int("role").Default(1).Comment("用户角色"),
		field.Time("created_at").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Time("updated_at").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
