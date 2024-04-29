package pokeapi

import (
	"net/http"
)

type LocationAreasResp struct {
	Previous *string `json:"previous"`
	Next     *string `json:"next"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
	Count int `json:"count"`
}

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreasResp, error) {
	fullUrl := baseUrl + "/location-area"
	if pageUrl != nil {
		fullUrl = *pageUrl
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	respData := LocationAreasResp{}
	err = c.fetchAndDecodeJson(req, &respData)
	if err != nil {
		return LocationAreasResp{}, err
	}

	return respData, nil
}
