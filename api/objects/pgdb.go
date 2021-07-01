package objects

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Pgdb struct {
	Host     string
	User     string
	Password string
	Dbname   string
	Port     int
	Sslmode  string
}

type PgConn struct {
	connection *gorm.DB
}

func (p Pgdb) Init() (PgConn, error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v", p.Host, p.User, p.Password, p.Dbname, p.Port, p.Sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return PgConn{connection: db}, err
}

func (p PgConn) CloseConnection() {
	db, _ := p.connection.DB()
	db.Close()
}

func (p PgConn) GetByOriginalUrl(url string) Golink {
	var query_result Golink
	p.connection.Where("original_url = ?", url).First(&query_result)
	return query_result
}

func (p PgConn) GetByGolink(link string) Golink {
	var query_result Golink
	p.connection.Where("golink = ?", link).First(&query_result)
	return query_result
}

func (p PgConn) GetAll() []Golink {
	var query_result []Golink
	p.connection.Find(&query_result)
	return query_result
}

func (p PgConn) CreateGolink(golink *Golink) {
	p.connection.Create(golink)
}
