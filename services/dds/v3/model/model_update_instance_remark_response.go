package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateInstanceRemarkResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateInstanceRemarkResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateInstanceRemarkResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceRemarkResponse", string(data)}, " ")
}
