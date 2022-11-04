package telegram

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

type Client struct {
	host    string
	baseUrl string
	client  http.Client
}

func New(host string, token string) Client {
	return Client{
		host:    host,
		baseUrl: newBaseUrl(token),
		client:  http.Client{},
	}
}

func newBaseUrl(token string) string {
	return "bot" + token
}

func (c *Client) Updates(offset int, limit int) ([]Update, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))
}

func (c *Client) doRequest(method string, query url.Values) ([]byte, error) {
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.baseUrl, method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("Не удалось выполнить запрос", err)
	}
	req.URL.RawQuery = query.Encode()
}

func (c *Client) SendMessage() {

}
