package pocket

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

func (c *Client) AddItem(r map[string]interface{}) error {
	p := url.Values{}
	p.Set("consumer_key", c.consumer_key)
	p.Set("access_token", c.access_token)

	p = requestOption(r, p)

	c.endpoint = "https://getpocket.com/v3/add"
	res, err := c.c.PostForm(c.endpoint, p)

	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New(fmt.Sprintf("HTTP request failed with %s", res.Status))
	}

	defer res.Body.Close()

	var pocket Pocket

	err = json.NewDecoder(res.Body).Decode(&pocket)
	if pocket.Status != 1 {
		return errors.New(fmt.Sprint("error add item"))
	}
	return nil
}
