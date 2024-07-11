package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// NOTE:GO HTTP GET request
func (c *Client) ListLocationAreas() (LocationAreaResponse, error) {
	endpoint := "/location-area/"
	fullURL := baseUrl + endpoint
	fmt.Printf("my url: %s\n", fullURL)

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	if res.StatusCode > 399 {
		return LocationAreaResponse{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	locationAreasResp := LocationAreaResponse{}
	err = json.Unmarshal(data, &locationAreasResp)

	if err != nil {
		return LocationAreaResponse{}, err
	}

	return locationAreasResp, nil
}
