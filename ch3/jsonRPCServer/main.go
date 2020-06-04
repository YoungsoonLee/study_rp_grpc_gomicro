package main

import (
	jsonparse "encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/rpc"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc/json"
)

type Args struct {
	ID string
}

type Book struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Author string `json:"author,omitempty"`
}

type JSONServer struct{}

func (t *JSONServer) GiveBookDetail(r *http.Request, args *Args, reply *Book) error {
	var books []Book

	absPath, err := filepath.Abs("ch3/jsonRPCServer/books.json")
	if err != nil {
		log.Fatal(err)
	}

	raw, readerr := ioutil.ReadFile(absPath)
	if readerr != nil {
		log.Println("error: ", readerr)
		os.Exit(1)
	}

	marshalerr := jsonparse.Unmarshal(raw, &books)
	if marshalerr != nil {
		log.Println("error: ", marshalerr)
		os.Exit(1)
	}

	for _, book := range books {
		if book.ID == args.ID {
			*reply = book
			break
		}
	}
	return nil

}

func main() {
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterService(new(JSONServer), "")

	r := mux.NewRouter()
	r.Handle("/rpc", s)

	http.ListenAndServe(":1234", r)
}
