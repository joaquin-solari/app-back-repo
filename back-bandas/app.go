package main

import (
    "fmt"
    "math/rand"
    "net/http"
    "time"
)

func main() {
    // Inicializamos la fuente de números aleatorios con una semilla única
    rand.Seed(time.Now().UnixNano())

    // Lista de nombres de bandas de rock
    bandasDeRock := []string{
        "U2",
        "Arctic Monkeys",
        "Radiohead",
        "The Strokes",
        "RHCP",
        "Muse",
        "Oasis",
        "Airbag",
        "Catfish and the bottlemen",
        "The kooks",
    }

    // Definimos el handler para el GET request en el root
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Elegimos una banda de rock aleatoria de la lista
        bandaAleatoria := bandasDeRock[rand.Intn(len(bandasDeRock))]
        fmt.Fprintf(w, "Banda de rock aleatoria: %s", bandaAleatoria)
    })

    // Iniciamos el servidor en el puerto 8080
    fmt.Println("Servidor escuchando en el puerto 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}
