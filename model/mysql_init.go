package model

import (
	"github.com/ingirls/common/mysql"
	proto "github.com/ingirls/user/proto/user"
)

// AutoMigrate AutoMigrate
func AutoMigrate(db ...string) {
	name := "default"
	if len(db) > 0 {
		name = db[0]
	}
	mysql.With(name).Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARACTER SET=utf8mb4").AutoMigrate(&proto.User{})
	mysql.With(name).Model(&proto.User{}).AddIndex("index_email", "email")
	mysql.With(name).Model(&proto.User{}).AddUniqueIndex("unique_email", "email")
	mysql.With(name).Model(&proto.User{}).AddIndex("index_username", "username")
	mysql.With(name).Model(&proto.User{}).AddUniqueIndex("unique_username", "username")
	mysql.With(name).Model(&proto.User{}).AddIndex("index_mobile", "mobile")
	mysql.With(name).Model(&proto.User{}).AddUniqueIndex("unique_mobile", "mobile")
}
