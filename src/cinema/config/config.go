package config

import(
    "os"
    "fmt"
    "io/ioutil"
    "encoding/json"
    "github.com/jmoiron/sqlx"
  _ "github.com/go-sql-driver/mysql"
)

// AppPort stores port configuration for the server
var AppPort string = SetConfigValue("GOPORT", "7000")

// DomainName is the domain name will allow to connect with the Dal
var DomainName string = SetConfigValue("DOMAINNAME", "https://localhost")

// DB user
var User string = SetConfigValue("db_user", "root")

// DB password
var Password string = SetConfigValue("db_password", "root")

// DB host
var Host string = SetConfigValue("db_host", "127.0.0.1:3306")

// Database name
var Database string = SetConfigValue("db", "cinema_tkt")

// Connection type
var Type string = SetConfigValue("TYPE", "mysql")

// Connection string
var Connectiondetails string = "" + User + ":" + Password + "@tcp(" + Host + ")/" + Database + ""

// SMTP password
var SmtpPassword string = SetConfigValue("SmtpPassword", "password")

// Mail from
var MailFrom string = SetConfigValue("MailFrom", "from")

// SMTP Host
var SmtpHost string = SetConfigValue("SmtpHost", "SmtpHost")

// Smtp Port
var SmtpPort string = SetConfigValue("SmtpPort", "port")

// SetConfigValue set environment by value and string
func SetConfigValue(e string, d string) string {
    jsonFile, err := os.Open("./config.json")
    if err != nil {
        fmt.Println(err)
    }
    defer jsonFile.Close()
    byteValue, _ := ioutil.ReadAll(jsonFile)
    var result map[string]interface{}
    json.Unmarshal([]byte(byteValue), &result)

    r := result[e]
    if r != nil {
        if v := result[e].(string); v != ""{
            return string(v)
        }
    }

    return d
}

// ConnectDB creates a connection and ping to the specific database
func ConnectDB() (*sqlx.DB, error) {

    return sqlx.Connect(Type, Connectiondetails)
}
