package domains

type Webhook struct {
	Method    string `form:"method"`
	Url       string `form:"url"`
	Timestamp string `form:"timestamp"`
	Content   any    `form:"content,omitempty"`
}

type WebhookUsecase interface {
	Store(w *Webhook) (statuscode int, err error)
	Get() (w []Webhook, statuscode int, err error)
}
