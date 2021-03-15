package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type AddImageTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddImageTagResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AddImageTagResponse struct{}"
	}

	return strings.Join([]string{"AddImageTagResponse", string(data)}, " ")
}
