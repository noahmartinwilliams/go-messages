package main

import "net/http"
import "time"
import "log"
import "io"

func main() {
	dur := 48 * time.Hour
	adderc := make(chan Message, 100)
	requestcc := make(chan chan Message, 100)
	MessageJar(dur, adderc, requestcc)

	adderc <- Message{Name:"system", Contents:"First message"}

	handleComments := func( w http.ResponseWriter, req *http.Request) {
		adderc <- Message{Name:req.FormValue("name"), Contents:req.FormValue("contents")}
		requestc := make(chan Message, 100)
		requestcc <- requestc
		str := MsgsToHTML(MsgToHTML(MsgEscape(requestc)))
		io.WriteString(w, str)
	}

	handleMessages := func( w http.ResponseWriter, req *http.Request) {
		requestc := make(chan Message, 100)
		requestcc <- requestc
		str := MsgsToHTML(MsgToHTML(MsgEscape(requestc)))
		io.WriteString(w, str)
	}

	http.Handle("/", http.RedirectHandler("/messages", 300))
	http.HandleFunc("/messages", handleMessages)
	http.HandleFunc("/comment", handleComments)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
