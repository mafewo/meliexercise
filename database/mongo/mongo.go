package mongo

import (
	"log"
	"time"

	"github.com/mafewo/meliexercise/config"

	"gopkg.in/mgo.v2"
)

// var md *mgo.Session

//ConnString es el struct que contiene el string DNS de conexion
type ConnString struct {
	conn string
}

//ConnMG tiene los datos de conexion
type ConnMG struct {
	user  string
	pass  string
	host  string
	port  string
	table string
	db    string
}

//NewMG retorna el string de conexion basado en el struct de conexion
func NewMG(user string, pass string, host string, port string, table string, db string) *ConnMG {
	return &ConnMG{
		user:  user,
		pass:  pass,
		host:  host,
		port:  port,
		table: table,
		db:    db,
	}
}

//InitializeDatabase abre la conexion mysql
func (c *ConnMG) InitializeDatabase() (*mgo.Session, *mgo.Collection, error) {
	log.Println("Consultando MongoDB ...")
	timeOut := time.Duration(3 * time.Second)
	mg, err := mgo.DialWithTimeout("mongodb://"+c.host, timeOut)
	if err != nil {
		return nil, nil, err
	}
	mg.SetPoolLimit(config.MgMaxConn)
	collection := mg.DB(c.db).C(c.table)
	return mg, collection, nil
}
