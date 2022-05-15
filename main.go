package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"time"
)

type Engine struct {
}

func (e *Engine) someHeavyComputing(c chan bool) {
	time.Sleep(1 * time.Minute)
	c <- true
}

var engine *Engine

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	c := make(chan bool)
	go engine.someHeavyComputing(c)
	// here some heavy computing too
	flag := <-c
	fmt.Println(flag)
}

func main() {
	fmt.Println("welcome")
	engine = &Engine{}
	s := http.Server{Addr: ":8000"}
	r := mux.NewRouter()
	r.HandleFunc("/welcome", welcomeHandler)
	s.Handler = r

	if err := s.ListenAndServe(); err != nil {
		fmt.Printf("error on ListenAndServe. err: %v", err)
		os.Exit(1)
	}
}
