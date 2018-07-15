package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-ini/ini"
)

// APPURL data
var APPURL string

// HTTP data
var HTTP bool

// HTTPS data
var HTTPS bool

// WS data
var WS bool

// LocalTimeZone string
var LocalTimeZone *time.Location

// Address data
var Address string

// IPFilter data
var IPFilter string

// IPList slice of IPs
var IPList []string

// Port data
var Port string

// SslPort data
var SslPort string

// WsAddr data
var WsAddr string

// Cert data
var Cert string

// Key data
var Key string

//JWTSecret data
var JWTSecret string

//Mongo data
var Mongo map[string]*DB

//MysqlMaxConn Max mysql Connections
var MysqlMaxConn int

//MysqlMaxIdleConn Max mysql Idle Connections
var MysqlMaxIdleConn int

//MgMaxConn Max mongo Connections
var MgMaxConn int

//Mysql data
var Mysql map[string]*DB

//Mail data
var Mail map[string]*MailCfg

//MailCfg struct for conection
type MailCfg struct {
	Sender   string `ini:"sender"`
	Account  string `ini:"account"`
	Password string `ini:"password"`
	Address  string `ini:"address"`
	Port     string `ini:"port"`
	Mode     string `ini:"mode"`
}

//CrmDB var
var CrmDB DB

//DB Struct data
type DB struct {
	User string `ini:"user"`
	Pass string `ini:"pass"`
	Host string `ini:"host"`
	Port string `ini:"port"`
	DB   string `ini:"db"`
}

// AllowOrigin CORS headers
var AllowOrigin []string

// AllowMethods CORS headers
var AllowMethods []string

// AllowHeaders CORS headers
var AllowHeaders []string

//Read initialize the config
func Read(pathToFile string) {

	/// Get INI
	cfg, err := ini.Load(pathToFile)
	if err != nil {
		fmt.Printf("Error al leer el archivo INI: %v", err)
		os.Exit(1)
	}

	APPURL = cfg.Section("paths").Key("url").MustString("")
	HTTP = cfg.Section("server").Key("http").MustBool(true)
	HTTPS = cfg.Section("server").Key("https").MustBool(false)
	WS = cfg.Section("server").Key("ws").MustBool(false)
	myLocation := cfg.Section("server").Key("localTimeZone").MustString("Local")
	LocalTimeZone, _ = time.LoadLocation(myLocation)
	log.Printf("Configuracion de hora local : %s =>", LocalTimeZone)

	Address = cfg.Section("server").Key("address").MustString("localhost")
	IPFilter = cfg.Section("server").Key("IPFilter").MustString("none")
	IPList = cfg.Section("server").Key("IPList").Strings(",")
	Port = strconv.Itoa(cfg.Section("server").Key("http_port").MustInt(8888))
	SslPort = strconv.Itoa(cfg.Section("server").Key("https_port").MustInt(443))
	WsAddr = strconv.Itoa(cfg.Section("server").Key("ws_port").MustInt(8000))

	Cert = cfg.Section("server").Key("cert").MustString("cert.pem")
	Key = cfg.Section("server").Key("key").MustString("key.pem")
	JWTSecret = cfg.Section("jwt").Key("secret").MustString("")
	MysqlMaxConn = cfg.Section("db.mysql.settings").Key("max").MustInt(10)
	MysqlMaxIdleConn = cfg.Section("db.mysql.settings").Key("idle").MustInt(0)
	MgMaxConn = cfg.Section("db.mongo.settings").Key("max").MustInt(10)

	mysqlMain := new(DB)
	cfg.Section("db.mysql.main").MapTo(mysqlMain)
	mysqlSec := new(DB)
	cfg.Section("db.mysql.sec").MapTo(mysqlSec)

	Mysql = map[string]*DB{
		"main": mysqlMain,
		"sec":  mysqlSec,
	}

	CrmDB.User = cfg.Section("MS_CRM").Key("user").MustString("")
	CrmDB.Pass = cfg.Section("MS_CRM").Key("pass").MustString("")
	CrmDB.Host = cfg.Section("MS_CRM").Key("host").MustString("")
	CrmDB.Port = cfg.Section("MS_CRM").Key("port").MustString("")
	CrmDB.DB = cfg.Section("MS_CRM").Key("db").MustString("")

	mongoMain := new(DB)
	cfg.Section("db.mongo.main").MapTo(mongoMain)

	Mongo = map[string]*DB{
		"main": mongoMain,
	}

	AllowOrigin = cfg.Section("cors").Key("allow-origin").Strings(",")
	AllowMethods = cfg.Section("cors").Key("allow-methods").Strings(",")
	AllowHeaders = cfg.Section("cors").Key("allow-headers").Strings(",")

	mailMain := new(MailCfg)
	cfg.Section("mail.presupuesto").MapTo(mailMain)
	Mail = map[string]*MailCfg{
		"main": mailMain,
	}
}
