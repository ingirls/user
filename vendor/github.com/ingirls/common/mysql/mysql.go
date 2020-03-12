package mysql

import (
	"fmt"

	"log"

	"github.com/jinzhu/gorm"

	// mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Conf struct
type Conf struct {
	User     string
	Password string
	Host     string
	Port     string
	Db       string
}

// MysqlDriver MysqlDriver
var MysqlDriver map[string]*gorm.DB

// With With
func With(name ...string) *gorm.DB {
	k := "default"
	if len(name) > 0 {
		k = name[0]
	}

	if _, ok := MysqlDriver[k]; ok {
		return MysqlDriver[k]
	}

	log.Fatalf("mysql [%s] not exist.", k)
	return nil
}

// InitMysql InitMysql
func InitMysql(confs map[string]Conf) {
	MysqlDriver = make(map[string]*gorm.DB)

	for name, conf := range confs {
		db, err := gorm.Open(
			"mysql",
			fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.User, conf.Password, conf.Host, conf.Port, conf.Db),
		)

		if err != nil {
			log.Fatalf("mysql [%s] connect fail. err [%s]", name, err)
			continue
		}

		log.Printf("mysql [%s] connected.", name)
		MysqlDriver[name] = db
	}
}
