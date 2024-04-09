package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"webhooks/internal/business/domains"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func GetDocuments[T any](client *firestore.Client, ctx context.Context, col string) ([]T, error) {
	var results []T
	iter := client.Collection(col).OrderBy("Timestamp", firestore.Desc).Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to fetch documents: %w", err)
		}

		var item T
		if err := doc.DataTo(&item); err != nil {
			log.Printf("Failed to decode document: %v", err)
			continue
		}

		results = append(results, item)
	}

	return results, nil
}

func UploadDocument(client *firestore.Client, ctx context.Context, col string, wh domains.Webhook) error {
	_, _, err := client.Collection(col).Add(ctx, wh)
	return err
}
