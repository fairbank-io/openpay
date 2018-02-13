package openpay

import (
	"encoding/json"
	"net/http"
	"path"
)

// Defines the public interface required to access available 'webhooks' methods
type WebhooksAPI interface {
	// https://www.openpay.mx/docs/api/#crear-un-webhook
	Create(wh *Webhook) error

	// https://www.openpay.mx/docs/api/#obtener-un-webhook
	Get(whID string) (*Webhook, error)

	// https://www.openpay.mx/docs/api/#listado-de-webhook
	List() ([]Webhook, error)

	// https://www.openpay.mx/docs/api/#eliminar-un-webhook
	Delete(whID string) error
}

type webhooksClient struct {
	c *Client
}

func (wc *webhooksClient) Create(wh *Webhook) error {
	b, err := wc.c.request(&requestOptions{
		endpoint: "webhooks",
		method:   http.MethodPost,
		data:     wh,
	})
	if err != nil {
		return err
	}

	json.Unmarshal(b, wh)
	return nil
}

func (wc *webhooksClient) Get(whID string) (*Webhook, error) {
	b, err := wc.c.request(&requestOptions{
		endpoint: path.Join("webhooks", whID),
		method:   http.MethodGet,
		data:     nil,
	})
	if err != nil {
		return nil, err
	}

	w := &Webhook{}
	json.Unmarshal(b, w)
	return w, nil
}

func (wc *webhooksClient) List() ([]Webhook, error) {
	b, err := wc.c.request(&requestOptions{
		endpoint: "webhooks",
		method:   http.MethodGet,
		data:     nil,
	})
	if err != nil {
		return nil, err
	}

	var list []Webhook
	json.Unmarshal(b, &list)
	return list, nil
}

func (wc *webhooksClient) Delete(whID string) error {
	_, err := wc.c.request(&requestOptions{
		endpoint: path.Join("webhooks", whID),
		method:   http.MethodDelete,
		data:     nil,
	})
	return err
}