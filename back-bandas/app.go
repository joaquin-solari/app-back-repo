package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	bandasDeRock := []string{
		"U2",
		"Arctic Monkeys",
		"Radiohead",
		"The Strokes",
		"RHCP",
		"Muse",
		"Oasis",
		"Airbag",
		"Catfish and the Bottlemen",
		"The Kooks",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Elegimos una banda de rock aleatoria de la lista
		bandaAleatoria := bandasDeRock[rand.Intn(len(bandasDeRock))]

		// Crear un mapa para la respuesta en formato JSON
		response := map[string]string{
			"banda": bandaAleatoria,
		}

		// Convertir el mapa en JSON y enviarlo como respuesta
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("Servidor escuchando en el puerto 3001...")
	if err := http.ListenAndServe(":3001", nil); err != nil {
		panic(err)
	}
}
