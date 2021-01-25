package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ModifyJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ModifyJobResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ModifyJobResponse struct{}"
	}

	return strings.Join([]string{"ModifyJobResponse", string(data)}, " ")
}
