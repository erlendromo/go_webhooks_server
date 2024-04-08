package usecases

import (
	"context"
	"net/http"
	"webhooks/internal/business/domains"

	"cloud.google.com/go/firestore"
)

// TODO use this in the application in connection with firestore and r.Context()

type webhookUsecase struct {
	Client *firestore.Client
	Ctx    context.Context
}

func NewWebhookUsecase(client *firestore.Client, ctx context.Context) domains.WebhookUsecase {
	return &webhookUsecase{
		Client: client,
		Ctx:    ctx,
	}
}

func (wUC *webhookUsecase) Store(w *domains.Webhook) (statuscode int, err error) {
	return http.StatusOK, nil
}

func (wUC *webhookUsecase) Get() (w []domains.Webhook, statuscode int, err error) {
	return []domains.Webhook{}, http.StatusOK, nil
}

func (wUC *webhookUsecase) GetByID(id string) (w domains.Webhook, statuscode int, err error) {
	return domains.Webhook{}, http.StatusOK, nil
}
