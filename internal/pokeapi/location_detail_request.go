package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (ResponseLocationDetail, error) {
	location := locationName
	if location == "" {
		return ResponseLocationDetail{}, errors.New("A location name is required")
	}
	fullURL := baseURL + "/location-area/" + location

	if val, ok := c.cache.Get(fullURL); ok {
		locationDetailFromCache := ResponseLocationDetail{}
		err := json.Unmarshal(val, &locationDetailFromCache)
		if err != nil {
			return ResponseLocationDetail{}, err
		}
		fmt.Println("Location retrieved from cache")
		return locationDetailFromCache, nil
	}
	fmt.Println("Retrieving location detail from HTTP")

	request, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return ResponseLocationDetail{}, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return ResponseLocationDetail{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)

	if response.StatusCode > 399 {
		return ResponseLocationDetail{}, fmt.Errorf("bad status code: %v", response.StatusCode)
	}
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return ResponseLocationDetail{}, err
	}
	LocationDetailResp := ResponseLocationDetail{}
	err = json.Unmarshal(data, &LocationDetailResp)
	if err != nil {
		return ResponseLocationDetail{}, err
	}
	c.cache.Add(fullURL, data)
	return LocationDetailResp, nil
}
