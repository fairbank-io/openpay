package openpay

import (
	"encoding/json"
	"net/http"
	"path"
)

// Defines the public interface required to access available 'customers' methods
type CustomersAPI interface {
	// https://www.openpay.mx/docs/api/#crear-un-nuevo-cliente
	Create(customer *Customer) error

	// https://www.openpay.mx/docs/api/#actualizar-un-cliente
	Update(customer *Customer) error

	// https://www.openpay.mx/docs/api/#obtener-un-cliente-existente
	Get(customerID string) (*Customer, error)

	// https://www.openpay.mx/docs/api/#listado-de-clientes
	List(req *CustomersListRequest) ([]Customer, error)

	// https://www.openpay.mx/docs/api/#eliminar-un-cliente
	Delete(customerID string) error

	// https://www.openpay.mx/docs/api/#crear-una-tarjeta
	AddCard(customerID string, card *Card) error

	// https://www.openpay.mx/docs/api/#obtener-una-tarjeta
	GetCard(customerID, cardID string) (*Card, error)

	// https://www.openpay.mx/docs/api/#listado-de-tarjetas
	ListCards(customerID string, req *ListRequest) ([]Card, error)

	// https://www.openpay.mx/docs/api/#eliminar-una-tarjeta
	DeleteCard(customerID, cardID string) error
}

type customersClient struct {
	c *Client
}

func (cu *customersClient) Create(customer *Customer) error {
	b, err := cu.c.request(&requestOptions{
		endpoint: "customers",
		method:   http.MethodPost,
		data:     customer,
	})
	if err != nil {
		return err
	}

	json.Unmarshal(b, customer)
	return nil
}

func (cu *customersClient) Update(customer *Customer) error {
	b, err := cu.c.request(&requestOptions{
		endpoint: path.Join("customers", customer.ID),
		method:   http.MethodPut,
		data:     customer,
	})
	if err != nil {
		return err
	}

	json.Unmarshal(b, customer)
	return nil
}

func (cu *customersClient) Get(customerID string) (*Customer, error) {
	b, err := cu.c.request(&requestOptions{
		endpoint: path.Join("customers", customerID),
		method:   http.MethodGet,
		data:     nil,
	})
	if err != nil {
		return nil, err
	}

	c := &Customer{}
	json.Unmarshal(b, c)
	return c, nil
}

func (cu *customersClient) List(req *CustomersListRequest) ([]Customer, error) {
	b, err := cu.c.request(&requestOptions{
		endpoint: "customers",
		method:   http.MethodGet,
		data:     req,
	})
	if err != nil {
		return nil, err
	}

	var list []Customer
	json.Unmarshal(b, &list)
	return list, nil
}

func (cu *customersClient) Delete(customerID string) error {
	_, err := cu.c.request(&requestOptions{
		endpoint: path.Join("customers", customerID),
		method:   http.MethodDelete,
		data:     nil,
	})
	return err
}

func (cu *customersClient) AddCard(customerID string, card *Card) error {
	b, err := cu.c.request(&requestOptions{
		endpoint: path.Join("customers", customerID, "cards"),
		method:   http.MethodPost,
		data:     card,
	})
	if err != nil {
		return err
	}

	json.Unmarshal(b, card)
	return nil
}

func (cu *customersClient) GetCard(customerID, cardID string) (*Card, error) {
	b, err := cu.c.request(&requestOptions{
		endpoint: path.Join("customers", customerID, "cards", cardID),
		method:   http.MethodGet,
		data:     nil,
	})
	if err != nil {
		return nil, err
	}

	c := &Card{}
	json.Unmarshal(b, c)
	return c, nil
}

func (cu *customersClient) ListCards(customerID string, req *ListRequest) ([]Card, error) {
	b, err := cu.c.request(&requestOptions{
		endpoint: path.Join("customers", customerID, "cards"),
		method:   http.MethodGet,
		data:     req,
	})
	if err != nil {
		return nil, err
	}

	var list []Card
	json.Unmarshal(b, &list)
	return list, nil
}

func (cu *customersClient) DeleteCard(customerID, cardID string) error {
	_, err := cu.c.request(&requestOptions{
		endpoint: path.Join("customers", customerID, "cards", cardID),
		method:   http.MethodDelete,
		data:     nil,
	})
	return err
}
