package domains

import "context"

type Webhook struct {
	Event     string `form:"event"`
	Url       string `form:"url"`
	Timestamp string `form:"timestamp"`
}

type WebhookUsecase interface {
	Store(ctx context.Context, w *Webhook) (statuscode int, err error)
	Get(ctx context.Context) (w []Webhook, statuscode int, err error)
}
