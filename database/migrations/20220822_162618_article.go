package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// Article_20220822_162618 DO NOT MODIFY
type Article_20220822_162618 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Article_20220822_162618{}
	m.Created = "20220822_162618"

	migration.Register("Article_20220822_162618", m)
}

// Up Run the migrations
func (m *Article_20220822_162618) Up() {
	m.CreateTable("article", "innodb", "utf8mb4")
	m.PriCol("id").SetAuto(true).SetDataType("int(11)").SetUnsigned(true)
	m.NewCol("title").SetDataType("varchar(50)").SetNullable(false)
	m.NewCol("introduction").SetDataType("varchar(200) COLLATE utf8_unicode_ci").SetNullable(false)
	m.NewCol("content").SetDataType("text COLLATE utf8_unicode_ci").SetNullable(false)
	m.NewCol("image").SetDataType("varchar(32)").SetDefault("''").SetNullable(false)
	m.NewCol("category_id").SetDataType("tinyint(3)").SetUnsigned(true).SetDefault("1").SetNullable(false)
	m.NewCol("status").SetDataType("tinyint(3)").SetUnsigned(true).SetDefault("1").SetNullable(false)
	m.NewCol("sort").SetDataType("tinyint(3)").SetUnsigned(true).SetDefault("1").SetNullable(false)
	m.NewCol("view_count").SetDataType("int(11)").SetUnsigned(true).SetDefault("0").SetNullable(false)
	m.NewCol("comment_count").SetDataType("int(11)").SetUnsigned(true).SetDefault("0").SetNullable(false)
	m.NewCol("like_count").SetDataType("int(11)").SetUnsigned(true).SetDefault("0").SetNullable(false)
	m.NewCol("created_at").SetDataType("datetime").SetDefault("CURRENT_TIMESTAMP")
	m.NewCol("updated_at").SetDataType("datetime").SetDefault("CURRENT_TIMESTAMP")
	m.NewCol("published_at").SetDataType("datetime")
	m.SQL(m.GetSQL())

}

// Down Reverse the migrations
func (m *Article_20220822_162618) Down() {
	m.SQL("DROP TABLE IF EXISTS article")
}
