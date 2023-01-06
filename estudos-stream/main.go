package main

import (
	"fmt"
	"net/http"
	"time"

	"gocdn"
)

func main() {
	http.HandleFunc("/videos/", func(w http.ResponseWriter, r *http.Request) {
		// Abre o arquivo
		file, err := os.Open("dados/1.mp4")
		if err != nil {
			// Trate o erro aqui
		}
		defer file.Close()

		// Lê o conteúdo do arquivo para um buffer
		fileData, err := ioutil.ReadAll(file)
		if err != nil {
			// Trate o erro aqui
		}

		// Sirva o conteúdo do arquivo usando o gocdn
		gocdn.ServeContent(w, r, "1.mp4", time.Now(), gocdn.CacheRoute{"1h"}, fileData)
	})
	http.ListenAndServe(":8000", nil)
}
