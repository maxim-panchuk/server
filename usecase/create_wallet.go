package usecase

import (
	"context"
	"io"
	"log"
	"net/http"
	"server/initdata"
	"server/tonclient"
	"server/wallet"
)

func CreateWalletHandler(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalf(err.Error())
	}

	data, err := initdata.Parse(string(b))
	if err != nil {
		log.Fatalf(err.Error())
	}

	address, err := tonclient.CreateWallet(context.Background())
	if err != nil {
		log.Fatalf(err.Error())
	}

	if err := wallet.Get().SaveWalletAddress(data.User.ID, address); err != nil {
		log.Fatalf(err.Error())
	}

	w.Write([]byte("wallet created"))
}
