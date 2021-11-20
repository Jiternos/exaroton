package exaroton

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// Session struct to represent a connection to the Exaroton API.
type Session struct {
	// General configurable settings.
	// Authentication token for this session
	Token string

	// The http client used for REST requests
	Client *http.Client

	// The user agent used for REST APIs
	UserAgent string
}

// Request struct for the REST API result
type Request struct {
	Success string `json:"success"`
	Errors  string `json:"error"`
	Data    string `json:"data"`
}

// Send a request to the Exaroton REST API with the given method and path.
func (s *Session) Request(method, url string, body io.Reader) (response []byte, err error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return
	}

	req.Header.Add("Authorization", s.Token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("User-Agent", s.UserAgent)

	resp, err := s.Client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	var tempresponse map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&tempresponse)
	if err != nil {
		return
	}

	if tempresponse["success"] == true {
		response, err = json.Marshal(tempresponse["data"])
		if err != nil {
			return
		}

	} else if tempresponse["success"] == false {
		err = errors.New(tempresponse["error"].(string))
	}

	return
}
