package main

import (
	"log"
	"flag"
	"os"
	"net/http"
	"text/template"
	"path/filepath"
	"sync"
	"github.com/uchiiii/trace"
)

type templateHandler struct{
	once sync.Once
	filename string
	templ *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func(){
		t.templ = template.Must(template.ParseFiles(filepath.Join("./templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func main(){
	//flag.String("var name", "default val", "description")
	var addr = flag.String("addr", ":8080", "The addr of the application.")
	flag.Parse()

	r := newRoom()
	r.tracer = trace.New(os.Stdout)

	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)

	go r.run()

	log.Println("Starting web server on", *addr)
	if err:= http.ListenAndServe(*addr, nil); err != nil{
		log.Fatal("ListenAndServe:", err)
	}
}