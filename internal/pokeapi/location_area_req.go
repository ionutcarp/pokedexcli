package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreaResp, error) {
	fullURL := baseURL + "/location-area"
	if pageURL != nil {
		fullURL = *pageURL
	}

	if val, ok := c.cache.Get(fullURL); ok {
		locationsFromCache := LocationAreaResp{}
		err := json.Unmarshal(val, &locationsFromCache)
		if err != nil {
			return LocationAreaResp{}, err
		}
		fmt.Println("cache hit")
		return locationsFromCache, nil
	}
	fmt.Println("cache miss")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResp{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	if resp.StatusCode > 399 {
		return LocationAreaResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}

	LocationAreasResp := LocationAreaResp{}
	err = json.Unmarshal(data, &LocationAreasResp)
	if err != nil {
		return LocationAreaResp{}, err
	}

	c.cache.Add(fullURL, data)
	return LocationAreasResp, nil
}
