package domains

type Webhook struct {
	ID      string `form:"id"`
	Url     string `form:"url"`
	Country string `form:"country"`
	Event   string `form:"event"`
}

type WebhookUsecase interface {
	Store(w *Webhook) (statuscode int, err error)
	Get() (w []Webhook, statuscode int, err error)
	GetByID(id string) (w Webhook, statuscode int, err error)
}
