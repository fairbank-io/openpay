package openpay

import "fmt"

// Represents service errors
// https://www.openpay.mx/docs/api/#errores
type APIError struct {
	// Valid values are: request, internal, gateway
	Category string `json:"category,omitempty"`

	// Numeric code for the error
	// https://www.openpay.mx/docs/api/#c-digos-de-error
	Code uint `json:"error_code,omitempty"`

	// HTTP request status code
	HTTPCode uint `json:"http_code,omitempty"`

	// General description
	Description string `json:"description,omitempty"`

	// Unique identifier for the request
	RequestID string `json:"request_id,omitempty"`

	// Contains potential fraud flags detected
	FraudRules []string `json:"fraud_rules,omitempty"`
}

// Returns a descriptive text representation
func (e *APIError) Error() string {
	return fmt.Sprintf("%d: %s - %s", e.Code, e.Category, e.Description)
}
