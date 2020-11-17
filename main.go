package main

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"html/template"
	"log"
	"net/http"
	"time"
	"yalantis_go_school/config"
	"yalantis_go_school/dao"
	"yalantis_go_school/mysql"
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
	mux := http.NewServeMux()
	handler := http.HandlerFunc(handleDefaultPage)
	mux.Handle("/", sessionMiddleware(handler))
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", conf.Server.Host, conf.Server.Port), mux)
	log.Fatal(err)
}

func sessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil || cookie == nil {
			ip := r.RemoteAddr
			session, err := dao.GetSessionByIp(db, ip)

			if err != nil {
				fmt.Println(err)
			}

			newToken := generateToken(ip)
			if session != nil && cookie != nil && cookie.Value != session.Token {
				session.Ip = ip
				session.Token = newToken
				dao.UpdateSession(db, session)
			} else {
				session = &dao.Session{
					Ip:    ip,
					Token: newToken,
				}
				dao.SaveSession(db, session)
			}

			http.SetCookie(w, &http.Cookie{
				Name:    "session",
				Value:   session.Token,
				Expires: time.Now().Add(120 * time.Second),
			})
		}

		next.ServeHTTP(w, r)
	})
}

func generateToken(key string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(key), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hash to store:", string(hash))

	return base64.StdEncoding.EncodeToString(hash)
}

func handleDefaultPage(res http.ResponseWriter, req *http.Request) {
	uniqueUsers := dao.GetSessionNumbers(db)
	fmt.Println("Unique users: ", uniqueUsers)
	data := struct {
		Counter int64
	}{
		Counter: uniqueUsers,
	}
	err := tpl.ExecuteTemplate(res, "main.gohtml", data)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}
