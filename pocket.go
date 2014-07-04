// A Go wapper for the pocket API
package pocket

import (
	"encoding/json"
	"errors"
	"fmt"
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
	Tag            string `json:"tag"`
}

func NewClient(consumer_key, access_token string) *Client {
	return &Client{
		consumer_key,
		access_token,
		EndPoint,
		http.DefaultClient,
	}
}

func (c *Client) PocketList(r map[string]interface{}) (ItemList, error) {
	p := url.Values{}
	p.Set("consumer_key", c.consumer_key)
	p.Set("access_token", c.access_token)

	p = requestOption(r, p)

	res, err := c.c.PostForm(c.endpoint, p)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("HTTP request failed with %s", res.Status))
	}

	defer res.Body.Close()

	var pocket Pocket

	err = json.NewDecoder(res.Body).Decode(&pocket)
	if len(pocket.List) == 0 {
		return nil, errors.New(fmt.Sprint("no response data"))
	}

	if err != nil {
		return nil, err
	}
	return pocket.List, nil
}

func requestOption(r map[string]interface{}, p url.Values) url.Values {

	for k, v := range r {
		switch k {
		case "State":
			p.Set("state", fmt.Sprint(v))
		case "Favorite":
			p.Set("favorite", fmt.Sprint(v))
		case "Tag":
			p.Set("tag", fmt.Sprint(v))
		case "ContentType":
			p.Set("contentType", fmt.Sprint(v))
		case "Sort":
			p.Set("sort", fmt.Sprint(v))
		case "DetailType":
			p.Set("detailType", fmt.Sprint(v))
		case "Search":
			p.Set("search", fmt.Sprint(v))
		case "Domain":
			p.Set("domain", fmt.Sprint(v))
		case "Since":
			p.Set("since", fmt.Sprint(v))
		case "Count":
			p.Set("count", fmt.Sprint(v))
		case "Offset":
			p.Set("offset", fmt.Sprint(v))
		}
	}
	return p
}
