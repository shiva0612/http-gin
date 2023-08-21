package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
)

func main() {
	ech := make(chan error, 1)
	// signal.Notify(ech, os.Interrupt)
	go ss(ech)
	defer sameplDefer("main ka defer")

	r := gin.Default()

	go r.Run(":8090")
	<-ech
	close(ech)
	fmt.Println("main DONE")
}

func ss(ech chan error) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	close(c)
	signal.Stop(c)
	ech <- fmt.Errorf("ext quit")
}

func not_working2() {
	defer sameplDefer("main ka defer")
	r := gin.Default()

	srv := &http.Server{
		Handler: r,
		Addr:    ":8090",
	}

	srv.ListenAndServe()
	fmt.Println("main DONE")

}
func not_working() {
	defer sameplDefer("main ka defer")
	r := gin.Default()

	r.Run(":8090")
	fmt.Println("main DONE")
}
func sameplDefer(str string) {
	fmt.Println(str)
}
