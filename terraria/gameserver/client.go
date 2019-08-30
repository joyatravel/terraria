package gameserver

import (
	"fmt"
	"net/http"

	"github.com/cockroachdb/errors"
)

// NewClient creates a new Client.
func NewClient(addr string, opts ...func(*Config)) Client {
	cfg := Config{
		HTTPClient: new(http.Client),
	}
	for _, opt := range opts {
		opt(&cfg)
	}
	return Client{
		baseURL: addr,
		httpc:   cfg.HTTPClient,
	}
}

type (
	// A Client can connect to the Terraria gameserver using its REST API.
	Client struct {
		baseURL string
		httpc   *http.Client
	}

	// A Config configures a Client.
	Config struct {
		HTTPClient *http.Client
	}
)

// Ping returns an error if the gameserver is unreachable.
func (c Client) Ping() error {
	res, err := c.httpc.Head(fmt.Sprintf("%s/status", c.baseURL))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.Newf("gameserver: bad response status '%d'", res.StatusCode)
	}
	return nil
}

// Get issues a GET request to a path of the Terraria gameserver.
func (c Client) Get(path string) (*http.Response, error) {
	return c.httpc.Get(c.baseURL + path)
}
