package schema

import (
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
		field.String("password").MaxLen(100).Sensitive().Comment("密码"),
		field.String("nick_name").MaxLen(20).Comment("昵称"),
		field.String("head_url").MaxLen(200).Comment("头像"),
		field.Time("birthday").Comment("生日").SchemaType(map[string]string{
			dialect.MySQL: "date",
		}),
		field.String("address").MaxLen(200).Comment("地址"),
		field.Text("desc").Comment("描述信息"),
		field.Int("gender").Comment("性别"),
		field.Int("role").Default(1).Comment("用户角色"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
