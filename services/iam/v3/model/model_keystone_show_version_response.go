package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneShowVersionResponse struct {
	Version        *Version `json:"version,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o KeystoneShowVersionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneShowVersionResponse struct{}"
	}

	return strings.Join([]string{"KeystoneShowVersionResponse", string(data)}, " ")
}
