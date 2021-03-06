package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneShowProjectResponse struct {
	Project        *ProjectResult `json:"project,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o KeystoneShowProjectResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneShowProjectResponse struct{}"
	}

	return strings.Join([]string{"KeystoneShowProjectResponse", string(data)}, " ")
}
