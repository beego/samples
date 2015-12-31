package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Post_20151227_175928 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Post_20151227_175928{}
	m.Created = "20151227_175928"
	migration.Register("Post_20151227_175928", m)
}

// Run the migrations
func (m *Post_20151227_175928) Up() {
	// use m.Sql("CREATE TABLE ...") to make schema update
	m.Sql("CREATE TABLE post(`id` int(11) NOT NULL AUTO_INCREMENT,`title` varchar(128) NOT NULL,`body` longtext  NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Post_20151227_175928) Down() {
	// use m.Sql("DROP TABLE ...") to reverse schema update
	m.Sql("DROP TABLE `post`")
}
