package usecases

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"webhooks/internal/business/domains"
	"webhooks/internal/constants"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

// TODO use this in the application in connection with firestore and r.Context()

type dbWebhook struct {
	Client     *firestore.Client
	Ctx        context.Context
	Collection string
}

func NewDBWebhook(client *firestore.Client, ctx context.Context) domains.WebhookUsecase {
	return &dbWebhook{
		Client:     client,
		Ctx:        ctx,
		Collection: constants.WEBHOOK_TRIGGERS_COLLECTION,
	}
}

func (wUC *dbWebhook) Store(w *domains.Webhook) (statuscode int, err error) {
	statuscode = http.StatusOK

	_, _, err = wUC.Client.Collection(wUC.Collection).Add(wUC.Ctx, &w)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Internal error: %w", err)
	}

	return
}

func (wUC *dbWebhook) Get() (whs []domains.Webhook, statuscode int, err error) {
	statuscode = http.StatusOK
	iter := wUC.Client.Collection(wUC.Collection).OrderBy(constants.TIMESTAMP, firestore.Desc).Limit(constants.FIRESTORE_REQUEST_LIMIT).Documents(wUC.Ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return nil, http.StatusInternalServerError, fmt.Errorf("%s: %w", constants.FAILED_FETCH_DOCS, err)
		}

		var item domains.Webhook
		if err := doc.DataTo(&item); err != nil {
			log.Printf("%s: %v", constants.FAILED_DECODE_DOC, err)
			continue
		}

		whs = append(whs, item)
	}

	return whs, http.StatusOK, err
}
