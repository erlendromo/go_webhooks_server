package domains

import "context"

type Webhook struct {
	Event     string `form:"event" firestore:"Event"`
	Url       string `form:"url" firestore:"Url"`
	Timestamp string `form:"timestamp" firestore:"Timestamp"`
}

type WebhookUsecase interface {
	Store(ctx context.Context, w *Webhook) (statuscode int, err error)
	Get(ctx context.Context) (w []Webhook, statuscode int, err error)
}
