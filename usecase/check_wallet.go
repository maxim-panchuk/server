package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"server/initdata"
	"server/tonclient"
	"server/wallet"
)

func CheckWalletHandler(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalf(err.Error())
	}

	data, err := initdata.Parse(string(b))
	if err != nil {
		log.Fatalln(err)
	}

	type GetWalletResp struct {
		WalletExists bool   `json:"walletExists"`
		Address      string `json:"address"`
		Balance      string `json:"balance"`
	}

	resp := &GetWalletResp{
		WalletExists: false,
	}

	result, err := json.Marshal(resp)
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()
	var address string
	address, err = wallet.GetStorageDB().GetAddressByUserID(ctx, data.User.ID)
	if err != nil {
		if errors.Is(err, wallet.ErrNoSuchUser) {
			address, err = createWallet(ctx, data.User.ID)
			if err != nil {
				log.Fatalln(err)
			}
		} else {
			log.Fatalln(err)
		}
	}

	balance, err := tonclient.GetBalance(ctx, address)
	if err != nil && !errors.Is(err, tonclient.ErrWalletIsNotActive) {
		log.Fatalln(err)
	}

	resp.Balance = balance
	resp.Address = address
	resp.WalletExists = true

	result, err = json.Marshal(resp)
	if err != nil {
		log.Fatalln(err)
	}

	w.Write(result)
}

func createWallet(ctx context.Context, userID int64) (string, error) {
	address, err := tonclient.CreateWallet(ctx)
	if err != nil {
		log.Fatalf(err.Error())
	}

	if err := wallet.GetStorageDB().SaveWalletAddress(ctx, userID, address); err != nil {
		log.Fatalf(err.Error())
	}

	return address, nil
}
