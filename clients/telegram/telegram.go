package telegram

import (
	"encoding/json"
	"github.com/Umk1nus/bot-go-for-tg/lib"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

const (
	getUpdatesMethod  = "getUpdates"
	sendMessageMethod = "sendMessages"
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

	data, err := c.doRequest(getUpdatesMethod, q)
	if err != nil {
		return nil, err
	}

	var res UpdatesResponse

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res.Result, nil
}

func (c *Client) doRequest(method string, query url.Values) (data []byte, err error) {
	defer func() {
		err = lib.ErrorValidate("Не получилось выполнить запрос", err)
	}()

	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.baseUrl, method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil

}

func (c *Client) SendMessage(chatID int, text string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", text)

	_, err := c.doRequest(sendMessageMethod, q)
	if err != nil {
		return lib.ErrorValidate("Не удалось отправить сообщение", err)
	}

	return err
}
