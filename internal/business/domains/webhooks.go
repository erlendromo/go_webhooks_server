package domains

import "context"

type Webhook struct {
	Method    string `form:"method"`
	Url       string `form:"url"`
	Timestamp string `form:"timestamp"`
	Content   any    `form:"content,omitempty"`
}

type WebhookUsecase interface {
	Store(ctx context.Context, w *Webhook) (statuscode int, err error)
	Get(ctx context.Context) (w []Webhook, statuscode int, err error)
}
