package openpay

import (
	"encoding/json"
	"net/http"
	"path"
)

// Defines the public interface required to access available 'charges' methods
type ChargesAPI interface {
	// https://www.openpay.mx/docs/api/#crear-una-tarjeta
	AddCard(card *Card) error

	// https://www.openpay.mx/docs/api/#obtener-un-cargo
	Get(txID string) (*Transaction, error)

	// https://www.openpay.mx/docs/api/#listado-de-cargos
	List(req *ChargesListRequest) ([]Transaction, error)

	// https://www.openpay.mx/docs/api/#cargo-en-tienda
	AtStore(charge *ChargeAtStore) (*Transaction, error)

	// https://www.openpay.mx/docs/api/#cargo-en-banco
	AtBank(charge *ChargeAtBank) (*Transaction, error)

	// https://www.openpay.mx/docs/api/#con-id-de-tarjeta-o-token
	WithCard(charge *ChargeWithStoredCard) (*Transaction, error)

	// https://www.openpay.mx/docs/api/#confirmar-un-cargo
	Capture(txID string, amount float32) (*Transaction, error)

	// https://www.openpay.mx/docs/api/#devolver-un-cargo
	Refund(txID string, amount float32, description string) (*Transaction, error)
}

type chargesClient struct {
	c *Client
}

func (cc *chargesClient) AddCard(card *Card) error {
	// Add the card at merchant level
	b, err := cc.c.request(&requestOptions{
		endpoint: "cards",
		method:   http.MethodPost,
		data:     card,
	})
	if err != nil {
		return err
	}

	json.Unmarshal(b, card)
	return nil
}

func (cc *chargesClient) Get(txID string) (*Transaction, error) {
	b, err := cc.c.request(&requestOptions{
		endpoint: path.Join("charges", txID),
		method:   http.MethodGet,
		data:     nil,
	})
	if err != nil {
		return nil, err
	}

	tx := &Transaction{}
	json.Unmarshal(b, tx)
	return tx, nil
}

func (cc *chargesClient) List(req *ChargesListRequest) ([]Transaction, error) {
	b, err := cc.c.request(&requestOptions{
		endpoint: "charges",
		method:   http.MethodGet,
		data:     req,
	})
	if err != nil {
		return nil, err
	}

	var list []Transaction
	json.Unmarshal(b, &list)
	return list, nil
}

func (cc *chargesClient) AtStore(charge *ChargeAtStore) (*Transaction, error) {
	b, err := cc.c.request(&requestOptions{
		endpoint: "charges",
		method:   http.MethodPost,
		data:     charge,
	})
	if err != nil {
		return nil, err
	}

	tx := &Transaction{}
	json.Unmarshal(b, tx)
	return tx, nil
}

func (cc *chargesClient) AtBank(charge *ChargeAtBank) (*Transaction, error) {
	b, err := cc.c.request(&requestOptions{
		endpoint: "charges",
		method:   http.MethodPost,
		data:     charge,
	})
	if err != nil {
		return nil, err
	}

	tx := &Transaction{}
	json.Unmarshal(b, tx)
	return tx, nil
}

func (cc *chargesClient) WithCard(charge *ChargeWithStoredCard) (*Transaction, error) {
	b, err := cc.c.request(&requestOptions{
		endpoint: "charges",
		method:   http.MethodPost,
		data:     charge,
	})
	if err != nil {
		return nil, err
	}

	tx := &Transaction{}
	json.Unmarshal(b, tx)
	return tx, nil
}

func (cc *chargesClient) Capture(txID string, amount float32) (*Transaction, error) {
	b, err := cc.c.request(&requestOptions{
		endpoint: path.Join("charges", txID, "capture"),
		method:   http.MethodPost,
		data:     map[string]float32{"amount":amount},
	})
	if err != nil {
		return nil, err
	}

	tx := &Transaction{}
	json.Unmarshal(b, tx)
	return tx, nil
}

func (cc *chargesClient) Refund(txID string, amount float32, description string) (*Transaction, error) {
	b, err := cc.c.request(&requestOptions{
		endpoint: path.Join("charges", txID, "refund"),
		method:   http.MethodPost,
		data:     map[string]interface{}{
			"amount": amount,
			"description": description,
		},
	})
	if err != nil {
		return nil, err
	}

	tx := &Transaction{}
	json.Unmarshal(b, tx)
	return tx, nil
}