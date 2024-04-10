package usecases

import (
	"context"
	"net/http"
	"webhooks/internal/business/domains"
	"webhooks/internal/constants"
	"webhooks/internal/datasources/database"

	"cloud.google.com/go/firestore"
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
	err = database.UploadDocument(wUC.Client, wUC.Ctx, wUC.Collection, *w)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, err
}

func (wUC *dbWebhook) Get() (whs []domains.Webhook, s int, err error) {
	whs, err = database.FetchDocuments[domains.Webhook](wUC.Client, wUC.Ctx, wUC.Collection)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return whs, http.StatusOK, err
}

// TODO Remove if not gonna be implemented
func (wUC *dbWebhook) GetByID(id string) (w domains.Webhook, statuscode int, err error) {
	return w, http.StatusOK, nil
}
