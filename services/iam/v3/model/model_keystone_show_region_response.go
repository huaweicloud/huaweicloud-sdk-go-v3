package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneShowRegionResponse struct {
	Region         *Region `json:"region,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o KeystoneShowRegionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneShowRegionResponse struct{}"
	}

	return strings.Join([]string{"KeystoneShowRegionResponse", string(data)}, " ")
}
