package database

import (
	"context"
	"errors"
	"fmt"
	"webhooks/internal/constants"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func FetchDocuments(ctx context.Context, q firestore.Query) ([]*firestore.DocumentSnapshot, error) {
	iter := q.Documents(ctx)
	defer iter.Stop()

	var snapshots []*firestore.DocumentSnapshot
	for {
		snapshot, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("%s: %w", constants.FAILED_FETCH_DOCS, err)
		}

		snapshots = append(snapshots, snapshot)
	}

	if len(snapshots) == 0 {
		return nil, fmt.Errorf(constants.NO_DOCS_FOUND)
	}

	return snapshots, nil
}

func DocToData[T any](s *firestore.DocumentSnapshot) (T, error) {
	var data T
	if err := s.DataTo(&data); err != nil {
		return data, err
	}

	return data, nil
}
