package registry

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Doc struct {
	Name     string             `json:"name"`
	DistTags map[string]string  `json:"dist-tags"`
	Versions map[string]Version `json:"versions"`
}

type Version struct {
	Name         string            `json:"name"`
	Dependencies map[string]string `json:"dependencies"`
}

type Registry struct {
	url    string
	client *http.Client
}

func (c *Client) Registry() *Registry {
	return &Registry{
		url: c.baseURL + "/registry/",
		client: &http.Client{
			Timeout: time.Second * 5,
		},
	}
}

func (r *Registry) GetDoc(id string) (*Doc, error) {
	resp, err := r.client.Get(r.url + id)
	if err != nil {
		return nil, fmt.Errorf("failed to http get %s, failed with err %s", r.url+id, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("expected status code 200, instead got %d", resp.StatusCode)
	}

	doc := &Doc{}
	err = json.NewDecoder(resp.Body).Decode(doc)
	if err != nil {
		return nil, fmt.Errorf("failed to decode server response, %s", err)
	}

	return doc, nil
}
