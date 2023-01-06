package main

import (
    "net/http"
    "os"
    "time"
)

func main() {
    http.HandleFunc("/video", func(w http.ResponseWriter, r *http.Request) {
        // Abra o arquivo MP4
        file, err := os.Open("fibonacciDisney.mp4")
        if err != nil {
            http.Error(w, "Erro ao abrir o arquivo", http.StatusInternalServerError)
            return
        }
        defer file.Close()

        // Defina o tipo de conteúdo como vídeo MP4
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Content-Type", "video/mp4")

        // Envie o arquivo para o cliente
        http.ServeContent(w, r, "video.mp4", time.Time{}, file)
    })

    http.ListenAndServe(":8000", nil)
}
