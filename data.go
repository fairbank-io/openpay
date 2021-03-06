package openpay

import "time"

// Represents a base charge to be executed
// https://www.openpay.mx/docs/api/#cargos
type Charge struct {
	// Valid values are: card, store, bank_account
	Method string `json:"method,omitempty"`

	// Amount to charge, with up to two decimal digits
	Amount float32 `json:"amount,omitempty"`

	// Valid values: MXN or USD
	Currency string `json:"currency,omitempty"`

	// Basic description for the origin of the charge
	Description string `json:"description,omitempty"`

	// Unique identifier for the order, should be unique for all transactions
	OrderID string `json:"order_id,omitempty"`

	// Customer information
	// Required when executing the charge from a commerce
	Customer Customer `json:"customer,omitempty"`

	// For 'redirect' payments, send a email with the payment form
	SendEmail bool `json:"send_email"`

	// For 'redirect' payments,
	RedirectURL string `json:"redirect_url,omitempty"`
}

// Charges executed as virtual POS
// https://www.openpay.mx/docs/api/#con-terminal-virtual
type ChargeWithVirtualPOS struct {
	Charge

	// Use false for this kind of charges
	Confirm bool `json:"confirm"`
}

// Charges executed using a previously stored card id or token
// https://www.openpay.mx/docs/api/#con-id-de-tarjeta-o-token
type ChargeWithStoredCard struct {
	Charge

	// ID or token of a previously stored card
	SourceID string `json:"source_id,omitempty"`

	// Card security code, required when using stored cards
	CVV2 string `json:"cvv2,omitempty"`

	// Device identifier generated by the fraud prevention tool
	DeviceSessionID string `json:"device_session_id,omitempty"`

	// Specify if the charge should be executed immediately or the funds just reserved
	// for later execution
	Capture bool `json:"capture"`

	// Use loyalty points for payment, valid values are: ONLY_POINTS, MIXED, NONE
	UseCardPoints string `json:"use_card_points,omitempty"`

	// Specify if 3d secure should be used
	Use3DSecure bool `json:"use_3d_secure"`

	// Applicable payment plan, if any
	PaymentPlan *PaymentPlan `json:"payment_plan,omitempty"`

	// Additional metadata for the transaction
	Metadata map[string]string `json:"metadata,omitempty"`
}

// Charges to be executed in a convenience store
// https://www.openpay.mx/docs/api/#cargo-en-tienda
type ChargeAtStore struct {
	Charge

	// Expiration date for the charge reference in UTC and ISO 8601 format
	DueDate time.Time `json:"due_date"`
}

// Charges to be executed with a bank transfer reference
// https://www.openpay.mx/docs/api/#cargo-en-banco
type ChargeAtBank struct {
	Charge

	// Expiration date for the charge reference in UTC and ISO 8601 format
	DueDate time.Time `json:"due_date"`
}

// Represents an executed transaction
// https://www.openpay.mx/docs/api/#objeto-transacci-n
type Transaction struct {
	// Unique identifier
	ID string `json:"id,omitempty"`

	// Authorization code generated by the processor
	Authorization string `json:"authorization,omitempty"`

	// Valid values are: fee, charge, payout, transfer
	TransactionType string `json:"transaction_type,omitempty"`

	// Affectation to the account: in, out
	OperationType string `json:"operation_type,omitempty"`

	// Unique identifier for the order, unique for all transactions
	OrderID string `json:"order_id,omitempty"`

	// Set when the transaction is tied to a specific customer
	CustomerID string `json:"customer_id,omitempty"`

	// Transaction value, with up to two decimal digits
	Amount float32 `json:"amount,omitempty"`

	// Valid values: MXN or USD
	Currency string `json:"currency,omitempty"`

	// Used method when executing the transaction
	Method string `json:"method,omitempty"`

	// UTC in ISO 8601 format
	CreationDate time.Time `json:"creation_date,omitempty"`

	// Current transaction status: completed, in_progress, failed
	// https://www.openpay.mx/docs/api/#objeto-transaction-status
	Status string `json:"status,omitempty"`

	// Set on 'failed' transactions
	ErrorMessage string `json:"error_message,omitempty"`

	// Basic description for the transaction
	Description string `json:"description,omitempty"`

	// Bank account used, if any
	BankAccount *BankAccount `json:"bank_account,omitempty"`

	// Card used, if any
	Card *Card `json:"card,omitempty"`

	// Card points used, if any
	CardPoints *CardPoints `json:"card_points,omitempty"`
}

// Basic address information
// https://www.openpay.mx/docs/api/#objeto-direcci-n
type Address struct {
	// Usually used to specify street and number, required
	Line1 string `json:"line1,omitempty"`

	// Usually used to specify building, suite, delegation
	Line2 string `json:"line2,omitempty"`

	// Usually used to specify areas, neighborhood
	Line3 string `json:"line3,omitempty"`

	// Required
	PostalCode string `json:"postal_code,omitempty"`

	// Required
	State string `json:"state,omitempty"`

	// Required
	City string `json:"city,omitempty"`

	// In ISO_3166-1 format
	CountryCode string `json:"country_code,omitempty"`
}

// Monthly installments charge
// https://www.openpay.mx/docs/api/#objeto-paymentplan
type PaymentPlan struct {
	// Number of monthly installments for the operation: 3, 6, 9, 12
	Payments string `json:"payments,omitempty"`
}

// Provides details for charges at convenience stores
// https://www.openpay.mx/docs/api/#objeto-store
type Store struct {
	// Charge reference
	Reference string `json:"reference,omitempty"`

	// To be scanned at the store
	BarcodeURL string `json:"barcode_url,omitempty"`

	// To be used at supported stores
	PaybinReference string `json:"paybin_reference,omitempty"`

	//  To be scanned at the store where paybin is supported
	BarcodePaybinURL string `json:"barcode_paybin_url,omitempty"`
}

// Details of used card points
// https://www.openpay.mx/docs/api/#objeto-cardpoints
type CardPoints struct {
	// Number of points used
	Used uint `json:"used"`

	// Number of points remaining in the card after the transaction
	Remaining uint `json:"remaining"`

	// Transaction amount payed for with points
	Amount float32 `json:"amount"`

	// Message to be displayed to the customer
	Caption string `json:"caption,omitempty"`
}

// Custom card chain
// https://www.openpay.mx/docs/api/#objeto-paynetchain
type PaynetChain struct {
	// Chain name
	Name string `json:"name"`

	// URL for the chain logo image
	Logo string `json:"logo"`

	// URL for the chain thumbnail image
	Thumb string `json:"thumb"`

	// Maximum amount valid for transactions on the chain
	MaxAmount float32 `json:"max_amount"`
}

// Georeferenced location
// https://www.openpay.mx/docs/api/#objeto-geolocation
type Geolocation struct {
	// Latitude
	Lat float32 `json:"lat"`

	// Longitude
	Lng float32 `json:"lng"`

	// Google maps identifier
	PlaceID string `json:"place_id,omitempty"`
}

// Individual customer information
// https://www.openpay.mx/docs/api/#clientes
type Customer struct {
	// Unique identifier
	ID string `json:"id,omitempty"`

	// Internally used unique identifier
	ExternalID string `json:"external_id,omitempty"`

	// If the customer is able to hold balance
	RequiresAccount bool `json:"required_account"`

	// Registration date in UTC and ISO 8601 format
	CreationDate time.Time `json:"creation_date,omitempty"`

	// Customer's first name
	Name string `json:"name,omitempty"`

	// Customer's last name
	LastName string `json:"last_name,omitempty"`

	// Contact email address
	Email string `json:"email,omitempty"`

	// Contact phone number
	PhoneNumber string `json:"phone_number,omitempty"`

	// Current customer registry status, valid values are: active, deleted
	Status string `json:"status,omitempty"`

	// Current customer balance, up to two decimal digits
	Balance float32 `json:"balance"`

	// Special code to operate transactions with any bank in Mexico
	Clabe string `json:"clabe,omitempty"`

	// Postal address
	Address Address `json:"address,omitempty"`

	// Reference for payments in convenience stores
	Store Store `json:"store,omitempty"`
}

// Card representation
// https://www.openpay.mx/docs/api/#tarjetas
type Card struct {
	// Unique identifier
	ID string `json:"id,omitempty"`

	// Registration date in UTC and ISO 8601 format
	CreationDate time.Time `json:"creation_date,omitempty"`

	// Full name of the card's holder
	HolderName string `json:"holder_name,omitempty"`

	// Full card number, 16 or 19 digits
	CardNumber string `json:"card_number,omitempty"`

	// Card security code
	CVV2 string `json:"cvv2,omitempty"`

	// From the card's expiration date, 2 digits
	ExpirationMonth string `json:"expiration_month"`

	// From the card's expiration date, 2 digits
	ExpirationYear string `json:"expiration_year"`

	// Card's holder postal address
	Address Address `json:"address,omitempty"`

	// Charges can be performed with the card
	AllowsCharges bool `json:"allows_charges"`

	// Payments can be send with the card
	AllowsPayouts bool `json:"allows_payouts"`

	// Specify if this is a point based card
	PointsCard bool `json:"points_card"`

	// Available options are: visa, mastercard, carnet o american express
	Brand string `json:"brand,omitempty"`

	// Available options are: debit, credit, cash
	Type string `json:"type,omitempty"`

	// Issuer bank name
	BankName string `json:"bank_name,omitempty"`

	// Issuer bank code
	BankCode string `json:"bank_code,omitempty"`

	// Customer owner of the card, if any
	CustomerID string `json:"customer_id,omitempty"`

	// Device identifier generated by the fraud prevention tool, optional
	DeviceSessionID string `json:"device_session_id,omitempty"`
}

// Customer's bank account details
// https://www.openpay.mx/docs/api/#cuentas-bancarias
type BankAccount struct {
	// Unique identifier
	ID string `json:"id,omitempty"`

	// Registration date in UTC and ISO 8601 format
	CreationDate time.Time `json:"creation_date,omitempty"`

	// Friendly account identifier
	Alias string `json:"alias,omitempty"`

	// Full name of the card's holder
	HolderName string `json:"holder_name,omitempty"`

	// Special code to operate transactions with any bank in Mexico
	Clabe string `json:"clabe,omitempty"`

	// Issuer bank name
	BankName string `json:"bank_name,omitempty"`

	// Issuer bank code
	BankCode string `json:"bank_code,omitempty"`
}

// Request a paginated list of items
type ListRequest struct {
	// Maximum number of records
	Limit uint `json:"limit"`

	// Pagination offset
	Offset uint `json:"offset"`

	// Creation date in format 'yyyy-mm-dd'
	Creation string `json:"creation,omitempty"`

	// Creation date upper range in format 'yyyy-mm-dd'
	CreationGte string `json:"creation[gte],omitempty"`

	// Creation date lower range in format 'yyyy-mm-dd'
	CreationLte string `json:"creation[lte],omitempty"`
}

// Request a list of customers
// https://www.openpay.mx/docs/api/#listado-de-clientes
type CustomersListRequest struct {
	ListRequest

	// Internally used customer unique identifier
	ExternalID string `json:"external_id,omitempty"`
}

// Request a list of charges records
// https://www.openpay.mx/docs/api/#listado-de-cargos
type ChargesListRequest struct {
	// Amount to charge, with up to two decimal digits
	Amount float32 `json:"amount,omitempty"`

	// Amount upper range limit
	AmountGte string `json:"amount[gte],omitempty"`

	// Amount lower range limit
	AmountLte string `json:"amount[lte],omitempty"`

	// Valid values are:
	// IN_PROGRESS, COMPLETED, REFUNDED, CHARGEBACK_PENDING, CHARGEBACK_ACCEPTED,
	// CHARGEBACK_ADJUSTMENT, CHARGE_PENDING, CANCELLED, FAILED
	Status string `json:"status,omitempty"`

	// Charges related to a specific order
	OrderID string `json:"order_id,omitempty"`
}

// Webhook instance representation
// https://www.openpay.mx/docs/api/#objeto-webhook
type Webhook struct {
	// Unique identifier
	ID string `json:"id,omitempty"`

	// Webhook's endpoint
	URL string `json:"url,omitempty"`

	// Username value for basic credentials
	User string `json:"user,omitempty"`

	// Password value for basic credentials
	Password string `json:"password,omitempty"`

	// Current status of the instance, can be 'verified' or 'unverified'
	Status string `json:"status,omitempty"`

	// Valid events to be delivered for the instance
	// charge.refunded
	// charge.failed
	// charge.cancelled
	// charge.created
	// charge.succeeded
	// charge.rescored.to.decline
	// subscription.charge.failed
	// payout.created
	// payout.succeeded
	// payout.failed
	// transfer.succeeded
	// fee.succeeded
	// fee.refund.succeeded
	// spei.received
	// chargeback.created
	// chargeback.rejected
	// chargeback.accepted
	// order.created
	// order.activated
	// order.payment.received
	// order.completed
	// order.expired
	// order.cancelled
	// order.payment.cancelled
	EventTypes []string `json:"event_types,omitemtpy"`
}