package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const timeFormat = "2006-01-02T15:04:05.000Z07"

var client = &http.Client{}

func GetLiveEvents() (*LiveEvents, error) {

	req, err := http.NewRequest("GET", "https://esports-api.lolesports.com/persisted/gw/getLive?hl=en-US", nil)
	// ...
	req.Header.Add("x-api-key", `0TvQnueqKa5mxJntVWt0w4LpLfEkrV1Ta8rQBb9Z`)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Respuesta: %d", resp.StatusCode))
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte

	var result LiveEvents
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		return nil, err
	}

	return &result, nil
}

func GetDetails(id string, time *time.Time) ([]*Frame, error) {
	url := fmt.Sprintf("https://feed.lolesports.com/livestats/v1/details/%s", id)
	if time != nil {
		url = fmt.Sprintf("%s?startingTime=%s", url, time.Format(timeFormat))
	}

	req, err := http.NewRequest("GET", url, nil)
	// ...
	//req.Header.Add("x-api-key", `0TvQnueqKa5mxJntVWt0w4LpLfEkrV1Ta8rQBb9Z`)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Respuesta: %d", resp.StatusCode))
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte

	var result Frames
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		return nil, err
	}

	return result.Frames, nil
}

func GetWindow(id string, time *time.Time) (*Window, error) {
	if id == "108176672286663787" {
		fmt.Println("fafa")
	}

	url := fmt.Sprintf("https://feed.lolesports.com/livestats/v1/window/%s", id)
	if time != nil {
		url = fmt.Sprintf("%s?startingTime=%s", url, time.Format(timeFormat))
	}

	req, err := http.NewRequest("GET", url, nil)
	// ...
	//req.Header.Add("x-api-key", `0TvQnueqKa5mxJntVWt0w4LpLfEkrV1Ta8rQBb9Z`)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Respuesta: %d", resp.StatusCode))
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte

	var result Window
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		return nil, err
	}

	return &result, nil
}
