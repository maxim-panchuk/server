package tonclient

import (
	"context"
	"fmt"
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/ton"
	"github.com/xssnick/tonutils-go/ton/wallet"
	"log"
)

func GetBalance(ctx context.Context, clientAddress string) (string, error) {
	client := liteclient.NewConnectionPool()

	err := client.AddConnectionsFromConfigUrl(context.Background(),
		"https://ton.org/global.config.json")
	if err != nil {
		return "", fmt.Errorf("connection err: %w", err)
	}
	api := ton.NewAPIClient(client, ton.ProofCheckPolicyFast).WithRetry()
	ctx = client.StickyContext(ctx)
	b, err := api.CurrentMasterchainInfo(ctx)
	if err != nil {
		return "", fmt.Errorf("get block err: %w", err)
	}
	addr := address.MustParseAddr(clientAddress)

	res, err := api.WaitForBlock(b.SeqNo).GetAccount(ctx, b, addr)
	if err != nil {
		return "", fmt.Errorf("get account err: %w", err)
	}

	return res.State.Balance.String(), nil
}

func CreateWallet(ctx context.Context) (string, error) {
	client := liteclient.NewConnectionPool()
	cfg, err := liteclient.GetConfigFromUrl(context.Background(),
		"https://ton.org/global.config.json")
	if err != nil {
		return "", fmt.Errorf("get config err: %w", err)
	}

	err = client.AddConnectionsFromConfig(context.Background(), cfg)
	if err != nil {
		return "", fmt.Errorf("connection err: %w", err)
	}

	api := ton.NewAPIClient(client, ton.ProofCheckPolicyFast).WithRetry()
	api.SetTrustedBlockFromConfig(cfg)

	ctx = client.StickyContext(ctx)

	words := wallet.NewSeed()

	w, err := wallet.FromSeed(api, words, wallet.V4R2)
	if err != nil {
		return "", fmt.Errorf("FromSeed err: %w", err)
	}

	log.Println("wallet address:", w.WalletAddress())

	return w.WalletAddress().String(), nil
}
