package main

import (
	"net/http"
	"server/usecase"
)

const token = "7240774657:AAE6aygrHlXvunGiN19f3aKZ7FSxAkaCQ3g"

func main() {
	http.HandleFunc("/check-wallet",
		corsMiddleware(
			loggingMiddleware(
				authMiddleware(
					usecase.CheckWalletHandler))))

	http.ListenAndServe(":8080", nil)
}
