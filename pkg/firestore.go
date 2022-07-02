package pkg

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

type FirestoreService struct {
	client *firestore.Client
	ctx    context.Context
}

func NewFirestoreService() *FirestoreService {

	ctx := context.Background()

	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return &FirestoreService{
		client: client,
		ctx:    ctx,
	}
}

func (s *FirestoreService) GetUserRecord(docID string) (*firestore.DocumentSnapshot, error) {
	return s.client.Collection("records").Doc(docID).Get(s.ctx)
}

func (s *FirestoreService) DeleteRecord(docID, value string) error {
	record, err := s.GetUserRecord(docID)

	if err != nil {
		return err
	}

	cleanerData := record.Data()["data"].([]interface{})

	var newData []interface{}
	for _, r := range cleanerData {
		if r.(string) != value {
			newData = append(newData, r)
		}
	}

	_, err = s.client.Collection("records").Doc(docID).Set(s.ctx, map[string]interface{}{
		"data": newData,
	}, firestore.MergeAll)

	return err
}
func (s *FirestoreService) AddUserRecord(docID, value string) error {

	record, err := s.GetUserRecord(docID)

	if err != nil {
		return err
	}

	var cleanerData []interface{}
	if record.Data()["data"] != nil {
		cleanerData = record.Data()["data"].([]interface{})

		cleanerData = append(cleanerData, value)
	} else {
		cleanerData = append(cleanerData, value)
	}

	_, err = s.client.Collection("records").Doc(docID).Set(s.ctx, map[string]interface{}{
		"data": cleanerData,
	}, firestore.MergeAll)

	return err
}
func (s *FirestoreService) CloseClient() {
	s.client.Close()
}
