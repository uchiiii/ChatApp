package main

import (
	"flag"
	"fmt"
	"github.com/uchiiii/trace"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"
	//"github.com/stretchr/gomniauth/provider/facebook"
	"github.com/stretchr/objx"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/BurntSushi/toml"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("./templates", t.filename)))
	})
	data := map[string]interface{}{
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil{
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}
	t.templ.Execute(w, data)
}

func main() {
	var config Config
	_, err := toml.DecodeFile("./config.tml", &config)
	if err != nil {
		fmt.Println(err)
		return
	}
	//flag.String("var name", "default val", "description")
	var addr = flag.String("addr", ":8080", "The addr of the application.")
	flag.Parse()

	gomniauth.SetSecurityKey(config.Auth.SecurityKey)
	gomniauth.WithProviders(
		google.New(config.Auth.Each["google"].Id, config.Auth.Each["google"].Secret, config.Auth.Each["google"].RedirectURL),
		github.New(config.Auth.Each["github"].Id, config.Auth.Each["github"].Secret, config.Auth.Each["github"].RedirectURL),
	)

	r := newRoom()
	r.tracer = trace.New(os.Stdout)

	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)

	go r.run()

	log.Println("Starting web server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
