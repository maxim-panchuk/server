package usecase

import (
	"context"
	"encoding/json"
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
		w.Write([]byte("invalid init data"))
		return
	}

	type GetWalletResp struct {
		WalletExists bool   `json:"walletExists"`
		Address      string `json:"address"`
		Balance      string `json:"balance"`
	}

	resp := &GetWalletResp{
		WalletExists: false,
	}

	errResp, err := json.Marshal(resp)
	if err != nil {
		log.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
	}

	address, err := wallet.Get().GetAddressByUserID(data.User.ID)
	if err != nil {
		log.Println(err.Error())
		w.Write(errResp)
		return
	}

	balance, err := tonclient.GetBalance(context.Background(), address)
	if err != nil {
		log.Println(err.Error())
		w.Write(errResp)
		return
	}

	resp.WalletExists = true
	resp.Balance = balance
	resp.Address = address

	payload, err := json.Marshal(resp)
	if err != nil {
		log.Println(err.Error())
		w.Write(errResp)
		return
	}

	w.Write(payload)
	return
}
