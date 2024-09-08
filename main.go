package main

import (
	"io"
	"log"
	"net/http"
	"server/initdata"
	"time"
)

const token = "6535396266:AAEJxPBD1lWrWHDZb2hYK5paMeBmFcbP3U4"

func main() {
	http.HandleFunc("/wallet", func(w http.ResponseWriter, r *http.Request) {
		// Разрешение CORS
		w.Header().Set("Access-Control-Allow-Origin", "https://4aaa-88-201-232-88.ngrok-free.app")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, ngrok-skip-browser-warning")

		// Обработка запросов OPTIONS
		if r.Method == http.MethodOptions {
			return
		}

		b, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatalf(err.Error())
		}

		if err := initdata.Validate(string(b), token, time.Hour); err != nil {
			log.Fatalf(err.Error())
		}

		w.Write([]byte("validation successful"))
	})

	// Запуск сервера
	http.ListenAndServe(":8080", nil)
}
