package openpay

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"path"
	"time"
)

// API endpoints
const testAPI = "https://sandbox-api.openpay.mx/"
const liveAPI = "https://api.openpay.mx/"

// Main service handler
type Client struct {
	// Methods related to 'charges' management
	Charges ChargesAPI

	// Methods related to 'customers' management
	Customers CustomersAPI

	c           *http.Client
	key         string
	merchantID  string
	apiVersion  string
	userAgent   string
	apiEndpoint string
}

// Available configuration options, if not provided sane values will be
// used by default
type Options struct {
	// Time to wait for requests, in seconds
	Timeout uint

	// Time to maintain open the connection with the service, in seconds
	KeepAlive uint

	// Maximum network connections to keep open with the service
	MaxConnections uint

	// API version to use
	APIVersion string

	// User agent value to report to the service
	UserAgent string

	// Whether to use test or production environment
	UseProduction bool
}

// Network request options
type requestOptions struct {
	method   string
	endpoint string
	data     interface{}
}

// Return sane default configuration values
func defaultOptions() *Options {
	return &Options{
		Timeout:        30,
		KeepAlive:      600,
		MaxConnections: 100,
		APIVersion:     "v1",
		UserAgent:      "",
		UseProduction:  false,
	}
}

// NewClient will construct a usable service handler using the provided API key and
// configuration options, if 'nil' options are provided default sane values will
// be used
func NewClient(key, merchantID string, options *Options) (*Client, error) {
	if key == "" {
		return nil, errors.New("API key is required")
	}

	if merchantID == "" {
		return nil, errors.New("merchant ID is required")
	}

	// If no options are provided, use default sane values
	if options == nil {
		options = defaultOptions()
	}

	// Configure base HTTP transport
	t := &http.Transport{
		MaxIdleConns:        int(options.MaxConnections),
		MaxIdleConnsPerHost: int(options.MaxConnections),
		DialContext: (&net.Dialer{
			Timeout:   time.Duration(options.Timeout) * time.Second,
			KeepAlive: time.Duration(options.KeepAlive) * time.Second,
			DualStack: true,
		}).DialContext,
	}

	// Setup main client
	client := &Client{
		key:        key,
		merchantID: merchantID,
		apiVersion: options.APIVersion,
		userAgent:  options.UserAgent,
		c: &http.Client{
			Transport: t,
			Timeout:   time.Duration(options.Timeout) * time.Second,
		},
	}

	// Set client endpoint
	if options.UseProduction {
		client.apiEndpoint = liveAPI
	} else {
		client.apiEndpoint = testAPI
	}

	client.Charges = &chargesClient{c: client}
	client.Customers = &customersClient{c: client}
	return client, nil
}

// Dispatch a network request to the service
func (i *Client) request(r *requestOptions) ([]byte, error) {
	// Get request endpoint
	endpoint := i.apiEndpoint + path.Join(i.apiVersion, i.merchantID, r.endpoint)

	// Build request with headers and credentials
	data, _ := json.Marshal(r.data)
	req, _ := http.NewRequest(r.method, endpoint, bytes.NewReader(data))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(i.key, "")
	if i.userAgent != "" {
		req.Header.Add("User-Agent", i.userAgent)
	}

	// Execute request
	res, err := i.c.Do(req)
	if res != nil {
		// Properly discard request content to be able to reuse the connection
		defer io.Copy(ioutil.Discard, res.Body)
		defer res.Body.Close()
	}

	// Network level errors
	if err != nil {
		return nil, err
	}

	// Get response contents
	body, err := ioutil.ReadAll(res.Body)

	// Application level errors
	if res.StatusCode >= 400 {
		e := &APIError{}
		json.Unmarshal(body, e)
		return nil, e
	}
	return body, nil
}
