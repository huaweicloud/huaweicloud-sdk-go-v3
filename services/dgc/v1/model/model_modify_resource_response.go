package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ModifyResourceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ModifyResourceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ModifyResourceResponse struct{}"
	}

	return strings.Join([]string{"ModifyResourceResponse", string(data)}, " ")
}
