package main //pt 2 13.06

import (
	"fmt"
	"log"
)

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", store)
	server := NewAPIServer(":3000", store)
	server.Run()
	//fmt.Println("running")
}
