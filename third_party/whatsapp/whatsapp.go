package whatsapp

import "github.com/mrrizkin/crontastic/app/config"

type Whatsapp struct {
	token    string
	base_url string
}

func New(config *config.Config) *Whatsapp {
	token := config.WHATSAPP_SENDER_TOKEN
	base_url := config.WHATSAPP_SERVICE_URL

	return &Whatsapp{token, base_url}
}

func (w *Whatsapp) SendMessage(message string) error {

	// send message
	return nil
}
