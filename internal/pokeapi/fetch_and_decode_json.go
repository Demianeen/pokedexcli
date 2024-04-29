package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) fetchAndDecodeJson(request *http.Request, target interface{}) error {
	var data []byte
	cacheKey := request.URL.String()
	if cachedData, ok := c.cache.Get(cacheKey); ok {
		data = cachedData
	} else {
		resp, err := c.httpClient.Do(request)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode > 399 {
			return fmt.Errorf("bad status: %v", resp.StatusCode)
		}

		fetchedData, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		data = fetchedData
		c.cache.Add(cacheKey, fetchedData)
	}

	err := json.Unmarshal(data, &target)
	if err != nil {
		return err
	}
	return nil
}
