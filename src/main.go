package main

import (
	"fmt"
	"net/http"
)

type MessageHandler struct{ message string }

func (m MessageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.message)
}

// --------------------------------------------------

// func Welcome(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to this shit")
// }

// ==============================================================================
// func messageHandler(message string) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, message)
// 	})
// }

// func main() {
// 	mux := http.NewServeMux()
// 	fs := http.FileServer(http.Dir("static"))

//		mux.Handle("/", fs)
//		mux.Handle("/message", messageHandler("Fuck you"))
//		http.ListenAndServe(":8080", mux)
//	}
//
// ==============================================================================
// func main() {
// 	// created a new multiplexer
// 	mux := http.NewServeMux()

// 	// registered new handler to multiplexer
// 	mux.Handle("/", http.FileServer(http.Dir("static")))
// 	mux.Handle("/welcome", MessageHandler{message: "Welcome "})
// 	mux.Handle("/notwelcome", MessageHandler{message: "Fuck you"})

// 	// creted a new server with a address and Handler
// 	server := &http.Server{
// 		Addr:    ":8080",
// 		Handler: mux, // ServeMux "is a" handler ( ServeMux type satisfis 'Handler' interface )
// 	}
// 	server.ListenAndServe()
// }

// ======================= using gorilla mux ===================

//	func main() {
//		note := Note{"text/templates", "Template generates textual output"}
//		t := template.New("pp")
//		t, err := t.Parse(tmpl)
//		if err != nil {
//			log.Fatal("Parse :", err)
//		}
//		if err := t.Execute(os.Stdout, note); err != nil {
//			log.Fatal("Execute : ", err)
//			return
//		}
//	}
//
