package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type cliente struct {
	Nome     string `json:"nome"`
	Endereco string `json:"endereco"`
}

var clientes []cliente

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /cliente", createCliente)
	mux.HandleFunc("GET /cliente", getClientes)
	mux.HandleFunc("GET /cliente/{id}", getCliente)
	mux.HandleFunc("DELETE /cliente/{id}", deleteCliente)

	fmt.Println("running on 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func createCliente(w http.ResponseWriter, r *http.Request) {
	response, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	var c cliente
	if err := json.Unmarshal(response, &c); err != nil {
		w.WriteHeader(422)
		return
	}
	log.Println("Creating client")
	clientes = append(clientes, c)
	w.WriteHeader(200)
}

func getClientes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(clientes)
	log.Println("Getting clients")
}

func getCliente(w http.ResponseWriter, r *http.Request) {
	idAsString := r.PathValue("id")
	id, err := strconv.Atoi(idAsString)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	if id >= len(clientes) || id < 0 {
		w.WriteHeader(404)
		return
	}
	log.Println("Getting client", idAsString)
	json.NewEncoder(w).Encode(clientes[id])
	w.WriteHeader(200)
}

func deleteCliente(w http.ResponseWriter, r *http.Request) {
	idAsString := r.PathValue("id")
	id, err := strconv.Atoi(idAsString)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	if id >= len(clientes) || id < 0 {
		w.WriteHeader(404)
		return
	}
	log.Println("Deleting client", idAsString)
	clientes = append(clientes[:id], clientes[id+1:]...)
	w.WriteHeader(200)
}
