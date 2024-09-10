package wallet

import "sync"

type Storage interface {
	GetAddressByUserID(id int64) (string, error)
	SaveWalletAddress(id int64, address string) error
}

type StorageOM struct {
	db map[int64]string
}

func (s *StorageOM) GetAddressByUserID(id int64) (string, error) {
	v, ok := s.db[id]
	if !ok {
		return "", ErrNoSuchUser
	}
	return v, nil
}

func (s *StorageOM) SaveWalletAddress(id int64, address string) error {
	if _, ok := s.db[id]; ok {
		return ErrWalletAlreadyExists
	}
	s.db[id] = address
	return nil
}

var (
	singleton Storage
	once      sync.Once
)

func Get() Storage {
	once.Do(func() {
		singleton = &StorageOM{
			db: map[int64]string{
				715219007: "UQDLTJygXw37n7upvx2nP3LPmims2cwjqR3XR9V75zSazoTN",
			},
		}

	})

	return singleton
}
