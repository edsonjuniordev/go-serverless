package infra

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type SqlClient struct {
	Db *gorm.DB
}

// Função para buscar sempre uma única instância da conexão do banco
func GetConn() *gorm.DB {
	if db == nil {
		db, err := gorm.Open(mysql.Open("root:root@tcp(mycontainername:3306)/mydatabase"))
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		fmt.Println("Connected with the database")

		return db
	}

	return db
}

func NewSqlClient() *SqlClient {
	return &SqlClient{
		Db: GetConn(),
	}
}

// Função para fechar a conexão com o banco
func (s *SqlClient) Close() error {
	db, err := s.Db.DB()
	if err != nil {
		fmt.Println("Error closing database: ", err.Error())
		return err
	}
	err = db.Close()
	if err != nil {
		fmt.Println("Error closing database: ", err.Error())
		return err
	}
	return nil
}
