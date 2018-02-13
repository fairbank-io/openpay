package openpay

import (
	"fmt"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	// API key is required
	_, err := NewClient("", "", nil)
	if err == nil {
		t.Error("failed to detect missing API key")
	}

	// Use the test key and merchant id provided in the public documentation
	client, _ := NewClient("sk_e568c42a6c384b7ab02cd47d2e407cab", "mzdtln0bmtms6o3kck8f", nil)

	t.Run("Customers", func(t *testing.T) {
		testCustomer := &Customer{
			Name:            "Rick",
			LastName:        "Sanchez",
			Email:           "rick@mail.com",
			RequiresAccount: false,
			Address: Address{
				Line1:       "Calle 6 #910",
				City:        "Cordoba",
				State:       "VER",
				CountryCode: "MX",
				PostalCode:  "94560",
			},
		}

		t.Run("Create", func(t *testing.T) {
			err := client.Customers.Create(testCustomer)
			if err != nil {
				t.Error(err)
			}
		})

		t.Run("Update", func(t *testing.T) {
			testCustomer.PhoneNumber = "5544556677"
			err := client.Customers.Update(testCustomer)
			if err != nil {
				t.Error(err)
			}
		})

		t.Run("Get", func(t *testing.T) {
			c, err := client.Customers.Get(testCustomer.ID)
			if err != nil {
				t.Error(err)
			}
			if c.Email != testCustomer.Email {
				t.Error("invalid data received")
			}
		})

		t.Run("List", func(t *testing.T) {
			list, err := client.Customers.List(&CustomersListRequest{})
			if err != nil {
				t.Error(err)
			}
			if len(list) == 0 {
				t.Error("invalid data received")
			}
		})

		t.Run("Delete", func(t *testing.T) {
			err := client.Customers.Delete(testCustomer.ID)
			if err != nil {
				t.Error(err)
			}
		})
	})

	t.Run("Cards", func(t *testing.T) {
		testCustomer := &Customer{
			Name:            "Rick",
			LastName:        "Sanchez",
			Email:           "rick@mail.com",
			RequiresAccount: false,
			Address: Address{
				Line1:       "Calle 6 #910",
				City:        "Cordoba",
				State:       "VER",
				CountryCode: "MX",
				PostalCode:  "94560",
			},
		}
		client.Customers.Create(testCustomer)
		defer client.Customers.Delete(testCustomer.ID)
		card := &Card{
			HolderName:      fmt.Sprintf("%s %s", testCustomer.Name, testCustomer.LastName),
			CardNumber:      "4111111111111111",
			CVV2:            "401",
			ExpirationMonth: "10",
			ExpirationYear:  "19",
			Address:         testCustomer.Address,
		}

		t.Run("Add", func(t *testing.T) {
			card.CustomerID = testCustomer.ID
			err := client.Customers.AddCard(testCustomer.ID, card)
			if err != nil {
				t.Error(err)
			}
		})

		t.Run("Get", func(t *testing.T) {
			c2, err := client.Customers.GetCard(testCustomer.ID, card.ID)
			if err != nil {
				t.Error(err)
			}
			if c2.CardNumber != "411111XXXXXX1111" {
				t.Error("invalid data received")
			}
		})

		t.Run("List", func(t *testing.T) {
			list, err := client.Customers.ListCards(testCustomer.ID, &ListRequest{})
			if err != nil {
				t.Error(err)
			}
			if len(list) == 0 {
				t.Error("invalid data received")
			}
		})

		t.Run("Delete", func(t *testing.T) {
			err := client.Customers.DeleteCard(testCustomer.ID, card.ID)
			if err != nil {
				t.Error(err)
			}
		})
	})

	t.Run("BankAccounts", func(t *testing.T) {
		testCustomer := &Customer{
			Name:            "Rick",
			LastName:        "Sanchez",
			Email:           "rick@mail.com",
			RequiresAccount: false,
			Address: Address{
				Line1:       "Calle 6 #910",
				City:        "Cordoba",
				State:       "VER",
				CountryCode: "MX",
				PostalCode:  "94560",
			},
		}
		client.Customers.Create(testCustomer)
		defer client.Customers.Delete(testCustomer.ID)
		acc := &BankAccount{
			HolderName: "Juan Hernández Sánchez",
			Clabe: "012298026516924616",
		}

		t.Run("Add", func(t *testing.T) {
			err := client.Customers.AddBankAccount(testCustomer.ID, acc)
			if err != nil {
				t.Error(err)
			}
		})

		t.Run("Get", func(t *testing.T) {
			ba, err := client.Customers.GetBankAccount(testCustomer.ID, acc.ID)
			if err != nil {
				t.Error(err)
			}
			if ba.Clabe != "012XXXXXXXXXX24616" {
				t.Error("invalid data received")
			}
		})

		t.Run("List", func(t *testing.T) {
			list, err := client.Customers.ListBankAccounts(testCustomer.ID, &ListRequest{})
			if err != nil {
				t.Error(err)
			}
			if len(list) == 0 {
				t.Error("invalid data received")
			}
		})

		t.Run("Delete", func(t *testing.T) {
			err := client.Customers.DeleteBankAccount(testCustomer.ID, acc.ID)
			if err != nil {
				t.Error(err)
			}
		})
	})
	
	t.Run("Charges", func(t *testing.T) {
		// Create test customer
		testCustomer := &Customer{
			Name:            "Rick",
			LastName:        "Sanchez",
			Email:           "rick@mail.com",
			RequiresAccount: false,
			Address: Address{
				Line1:       "Calle 6 #910",
				City:        "Cordoba",
				State:       "VER",
				CountryCode: "MX",
				PostalCode:  "94560",
			},
		}
		client.Customers.Create(testCustomer)
		defer client.Customers.Delete(testCustomer.ID)

		// Test transaction ID holder
		txid := ""

		// Test card at merchant level
		card := &Card{
			HolderName:      fmt.Sprintf("%s %s", testCustomer.Name, testCustomer.LastName),
			CardNumber:      "4111111111111111",
			CVV2:            "401",
			ExpirationMonth: "10",
			ExpirationYear:  "19",
			Address:         testCustomer.Address,
		}

		t.Run("AddCard", func(t *testing.T) {
			err := client.Charges.AddCard(card)
			if err != nil {
				t.Error(err)
			}
		})

		t.Run("AtStore", func(t *testing.T) {
			_, err := client.Charges.AtStore(&ChargeAtStore{
				Charge: Charge{
					Method:      "store",
					Amount:      100,
					Currency:    "MXN",
					Description: "sample charge operation",
					Customer:    *testCustomer,
				},
				DueDate: time.Now().Add(72 * time.Hour),
			})
			if err != nil {
				t.Error(err)
			}
		})

		t.Run("AtBank", func(t *testing.T) {
			_, err := client.Charges.AtBank(&ChargeAtBank{
				Charge: Charge{
					Method:      "bank_account",
					Amount:      100,
					Currency:    "MXN",
					Description: "sample charge operation",
					Customer:    *testCustomer,
				},
				DueDate: time.Now().Add(72 * time.Hour),
			})
			if err != nil {
				t.Error(err)
			}
		})

		t.Run("Get", func(t *testing.T) {
			tx, err := client.Charges.Get(txid)
			if err != nil {
				t.Error(err)
			}
			if tx.ID != txid {
				t.Error("invalid data received")
			}
		})

		t.Run("List", func(t *testing.T) {
			list, err := client.Charges.List(&ChargesListRequest{})
			if err != nil {
				t.Error(err)
			}
			if len(list) == 0 {
				t.Error("invalid data received")
			}
		})

		t.Run("WithCard", func(t *testing.T) {
			tx, err := client.Charges.WithCard(&ChargeWithStoredCard{
				Charge: Charge{
					Method:      "card",
					Amount:      1000,
					Currency:    "MXN",
					Description: "sample charge operation",
					Customer:    *testCustomer,
				},
				SourceID:        card.ID,
				DeviceSessionID: card.DeviceSessionID,
				Capture:         false,
			})
			if err != nil {
				t.Error(err)
			}
			txid = tx.ID
		})

		t.Run("Capture", func(t *testing.T) {
			tx, err := client.Charges.Capture(txid, 1000)
			if err != nil {
				t.Error(err)
			}
			if tx.ID != txid {
				t.Error("invalid data received")
			}
		})

		t.Run("Refund", func(t *testing.T) {
			tx, err := client.Charges.Refund(txid, 1000, "refund sample operation")
			if err != nil {
				t.Error(err)
			}
			if tx.ID != txid {
				t.Error("invalid data received")
			}
		})
	})

	t.Run("Webhooks", func(t *testing.T) {
		hook := &Webhook{
			User: "foo",
			Password: "bar",
			URL: "https://hookb.in/voJJ3XXQ",
			EventTypes: []string{
				"charge.succeeded",
				"spei.received",
			},
		}

		t.Run("Create", func(t *testing.T) {
			err := client.Webhooks.Create(hook)
			if err != nil {
				t.Error(err)
			}
		})

		t.Run("Get", func(t *testing.T) {
			w2, err := client.Webhooks.Get(hook.ID)
			if err != nil {
				t.Error(err)
			}
			if w2.ID != hook.ID {
				t.Error("invalid data received")
			}
		})

		t.Run("List", func(t *testing.T) {
			list, err := client.Webhooks.List()
			if err != nil {
				t.Error(err)
			}
			if len(list) == 0 {
				t.Error("invalid data received")
			}
		})

		t.Run("Delete", func(t *testing.T) {
			err := client.Webhooks.Delete(hook.ID)
			if err != nil {
				t.Error(err)
			}
		})
	})
}
