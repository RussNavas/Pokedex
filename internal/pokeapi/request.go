package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error){
	/*
		ListLocationAreas fetches loaction areas from the PokeAPI.
		If pageURL is proved, it uses that URL; otherwise, 
		it defaults to the first page.
	*/

	url := "https://pokeapi.co/api/v2/location-area"
	if pageURL != nil{
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return LocationAreasResp{}, fmt.Errorf("bad status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	LocationsResp := LocationAreasResp{}
	err = json.Unmarshal(data, &LocationsResp)
	if err != nil {
		// why is this giving: LocationAreasResp is not a type
		return LocationAreasResp{}, err
	}

	return LocationsResp, nil
}
