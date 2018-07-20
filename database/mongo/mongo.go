package mongo

import (
	"crypto/tls"
	"log"
	"net"
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

//InitializeDatabase abre la conexion mongo
func (c *ConnMG) InitializeDatabase() (*mgo.Session, *mgo.Collection, error) {
	log.Println("Consultando MongoDB ...")
	dialInfo, err := mgo.ParseURL(c.host)
	if err != nil {
		log.Println(err)
	}
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		tlsConfig := &tls.Config{}
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		if err != nil {
			log.Println(err)
		}
		return conn, err
	}
	dialInfo.Timeout = time.Duration(3 * time.Second)
	//dialInfo.FailFast = true
	dialInfo.Username = c.user
	dialInfo.Password = c.pass
	mg, err := mgo.DialWithInfo(dialInfo)

	mg.SetPoolLimit(config.MgMaxConn)
	collection := mg.DB(c.db).C(c.table)
	return mg, collection, nil
}
