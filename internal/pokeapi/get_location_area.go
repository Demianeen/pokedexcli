package pokeapi

import (
	"net/http"
)

type LocationAreaResp struct {
	Encounter_method_rates []struct {
		Encounter_method struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"encounter_method"`
		Version_details []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Game_index int `json:"game_index"`
	Id         int `json:"id"`
	Location   struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	Pokemon_encounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"pokemon"`
		Version_details []struct {
			Encounter_details []struct {
				Chance           int           `json:"chance"`
				Condition_values []interface{} `json:"condition_values"`
				Max_level        int           `json:"max_level"`
				Method           struct {
					Name string `json:"name"`
					Url  string `json:"url"`
				} `json:"method"`
				Min_level int `json:"min_level"`
			} `json:"encounter_details"`
			Max_chance int `json:"max_chance"`
			Version    struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationAreaResp, error) {
	fullUrl := baseUrl + "/location-area/" + locationAreaName

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreaResp{}, err
	}

	respData := LocationAreaResp{}
	err = c.fetchAndDecodeJson(req, &respData)
	if err != nil {
		return LocationAreaResp{}, err
	}

	return respData, nil
}
