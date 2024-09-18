package wallet

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"log"
	"sync"
)

type Storage interface {
	GetAddressByUserID(context.Context, int64) (string, error)
	SaveWalletAddress(context.Context, int64, string) error
}

type StorageOM struct {
	db map[int64]string
}

func (s *StorageOM) GetAddressByUserID(_ context.Context, id int64) (string, error) {
	v, ok := s.db[id]
	if !ok {
		return "", ErrNoSuchUser
	}
	return v, nil
}

func (s *StorageOM) SaveWalletAddress(_ context.Context, id int64, address string) error {
	if _, ok := s.db[id]; ok {
		return ErrWalletAlreadyExists
	}
	s.db[id] = address
	return nil
}

type StorageDB struct {
	conn *pgx.Conn
}

func (s *StorageDB) GetAddressByUserID(ctx context.Context, id int64) (string, error) {
	var address string
	if err := s.conn.
		QueryRow(ctx, "SELECT address FROM twa_database.wallets WHERE user_id=$1", id).
		Scan(&address); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", ErrNoSuchUser
		}
		log.Fatalf("QueryRow failed: %v\n", err)
	}
	return address, nil
}

func (s *StorageDB) SaveWalletAddress(ctx context.Context, id int64, address string) error {
	if _, err := s.conn.
		Exec(ctx, "INSERT INTO twa_database.wallets (user_id, address) VALUES ($1, $2)", id, address); err != nil {
		log.Fatalf("Exec failed: %s\n", err.Error())
	}
	return nil
}

var (
	storageDBSingleton Storage
	storageDBOnce      sync.Once
)

func GetStorageDB() Storage {
	storageDBOnce.Do(func() {
		storageDBSingleton = &StorageDB{
			conn: connect(),
		}
	})
	return storageDBSingleton
}

var (
	storageOMSingleton Storage
	storageOMOnce      sync.Once
)

func GetStorageOM() Storage {
	storageOMOnce.Do(func() {
		storageOMSingleton = &StorageOM{
			db: make(map[int64]string),
		}

	})

	return storageOMSingleton
}

func connect() *pgx.Conn {
	url := "postgres://postgres:password@localhost:5432/postgres"
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	return conn
}
