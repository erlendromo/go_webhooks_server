package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"webhooks/internal/business/domains"
	"webhooks/internal/constants"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func FetchDocuments[T any](client *firestore.Client, ctx context.Context, col string) ([]T, error) {
	var results []T
	iter := client.Collection(col).OrderBy(constants.TIMESTAMP, firestore.Desc).Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("%s: %w", constants.FAILED_FETCH_DOCS, err)
		}

		var item T
		if err := doc.DataTo(&item); err != nil {
			log.Printf("%s: %v", constants.FAILED_DECODE_DOC, err)
			continue
		}

		results = append(results, item)
	}

	return results, nil
}

// TODO Modify structure to the same as above?
func UploadDocument(client *firestore.Client, ctx context.Context, col string, wh domains.Webhook) error {
	_, _, err := client.Collection(col).Add(ctx, wh)
	return err
}
