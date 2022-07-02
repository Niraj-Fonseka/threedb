package threedb

import (
	"errors"

	"cloud.google.com/go/firestore"
	"github.com/Niraj-Fonseka/threedb/pkg"
)

type ThreeDBInterface interface {
	Get(key string) (*firestore.DocumentSnapshot, error)
	Create(key string, value string) error
	Delete(key string, value string) error
	Update(key string) error
}

type Store struct {
	firestore *pkg.FirestoreService
}

func NewThreeDB(driver string) *Store {

	return &Store{
		firestore: pkg.NewFirestoreService(),
	}
}

func (s *Store) Get(key string) (*firestore.DocumentSnapshot, error) {
	return s.firestore.GetUserRecord(key)
}

func (s *Store) Create(key string, value string) error {
	return s.firestore.AddUserRecord(key, value)
}

func (s *Store) Delete(key, value string) error {
	return s.firestore.DeleteRecord(key, value)
}

//TODO
func (s *Store) Update(key string) error {
	return errors.New("update error")
}
