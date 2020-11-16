package main

import (
	"fmt"
	"gorm.io/gorm"
	"html/template"
	"log"
	"net/http"
	"yalantis_test/config"
	"yalantis_test/dao"
	"yalantis_test/mysql"
)

var tpl *template.Template
var db *gorm.DB
var conf config.Configuration

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
	conf = config.GetConfiguration()
	db = mysql.Open(conf)
}

func main() {
	http.HandleFunc("/", handleDefaultPage)
	http.ListenAndServe(fmt.Sprintf("%s:%s", conf.Server.Host, conf.Server.Port), nil)
}

func handleDefaultPage(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Header.Get("User-Agent"))
	data := struct {
		Counter int64
	}{
		Counter: dao.GetSessionNumbers(db),
	}
	err := tpl.ExecuteTemplate(res, "main.gohtml", data)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}