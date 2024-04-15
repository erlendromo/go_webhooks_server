package usecases

import (
	"context"
	"fmt"
	"net/http"
	"webhooks/internal/business/domains"
	"webhooks/internal/constants"
	"webhooks/internal/datasources/database"

	"cloud.google.com/go/firestore"
)

// TODO use this in the application in connection with firestore and r.Context()

type dbWebhook struct {
	Client     *firestore.Client
	Collection string
}

func NewDBWebhook(client *firestore.Client) domains.WebhookUsecase {
	return &dbWebhook{
		Client:     client,
		Collection: constants.WEBHOOK_TRIGGERS_COLLECTION,
	}
}

func (wUC *dbWebhook) Store(ctx context.Context, w *domains.Webhook) (statuscode int, err error) {
	statuscode = http.StatusOK

	_, _, err = wUC.Client.Collection(wUC.Collection).Add(ctx, &w)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("internal error: %w", err)
	}

	return
}

func (wUC *dbWebhook) Get(ctx context.Context) (whs []domains.Webhook, statuscode int, err error) {
	statuscode = http.StatusOK
	query := wUC.Client.Collection(wUC.Collection).OrderBy(constants.TIMESTAMP, firestore.Desc).Limit(constants.FIRESTORE_REQUEST_LIMIT)

	snapshots, err := database.FetchDocuments(ctx, query)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	for _, snapshot := range snapshots {
		w, err := database.DocToData[domains.Webhook](snapshot)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}

		whs = append(whs, w)
	}

	return
}
