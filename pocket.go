// A Go wapper for the pocket API
package pocket

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const EndPoint = "https://getpocket.com/v3/get"

type Client struct {
	consumer_key string
	access_token string
	endpoint     string
	c            *http.Client
}

type Pocket struct {
	Status   int      `json:"status"`
	Complete int      `json:"complete"`
	List     ItemList `json:"list"`
	Since    int      `json:"since"`
}

type ItemList map[string]Lists

type Lists struct {
	Item_id        string `json:"item_id"`
	Resolved_id    string `json:"resolved_id"`
	Given_url      string `json:"given_url"`
	Given_title    string `json:"given_title"`
	Favorite       string `json:"favorite"`
	Status         string `json:"status"`
	Time_added     string `json:"time_added"`
	Time_updated   string `json:"time_updated"`
	Time_favorited string `json:"fime_favorited"`
	Resolved_title string `json:"resolved_title"`
	Resolved_url   string `json:"resolved_url"`
	Excerpt        string `json:"excerpt"`
	Is_article     string `json:"is_article"`
	Is_index       string `json:"is_index"`
	Has_video      string `json:"has_video"`
	Has_image      string `json:"has_image"`
	Word_count     string `json:"word_count"`
}

func NewClient(consumer_key, access_token string) *Client {
	return &Client{
		consumer_key,
		access_token,
		EndPoint,
		http.DefaultClient,
	}
}

func (c *Client) PocketList() (ItemList, error) {
	p := url.Values{}
	p.Set("consumer_key", c.consumer_key)
	p.Set("access_token", c.access_token)

	res, err := c.c.PostForm(c.endpoint, p)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var pocket Pocket

	err = json.NewDecoder(res.Body).Decode(&pocket)
	if err != nil {
		return nil, err
	}
	return pocket.List, nil
}
