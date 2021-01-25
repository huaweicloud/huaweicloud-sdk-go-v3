package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateMpeCallBackResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateMpeCallBackResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateMpeCallBackResponse struct{}"
	}

	return strings.Join([]string{"CreateMpeCallBackResponse", string(data)}, " ")
}
