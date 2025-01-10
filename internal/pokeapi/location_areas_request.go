package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (ResponseLocationAreas, error) {
	fullURL := baseURL + "/location-area"
	if pageURL != nil {
		fullURL = *pageURL
	}

	if val, ok := c.cache.Get(fullURL); ok {
		locationsFromCache := ResponseLocationAreas{}
		err := json.Unmarshal(val, &locationsFromCache)
		if err != nil {
			return ResponseLocationAreas{}, err
		}
		fmt.Println("Locations list retrieved from cache")
		return locationsFromCache, nil
	}
	fmt.Println("Retrieving location areas from HTTP")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return ResponseLocationAreas{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ResponseLocationAreas{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	if resp.StatusCode > 399 {
		return ResponseLocationAreas{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return ResponseLocationAreas{}, err
	}

	LocationAreasResp := ResponseLocationAreas{}
	err = json.Unmarshal(data, &LocationAreasResp)
	if err != nil {
		return ResponseLocationAreas{}, err
	}

	c.cache.Add(fullURL, data)
	return LocationAreasResp, nil
}
