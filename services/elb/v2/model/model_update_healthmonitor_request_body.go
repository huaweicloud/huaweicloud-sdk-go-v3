package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type UpdateHealthmonitorRequestBody struct {
	Healthmonitor *UpdateHealthmonitorReq `json:"healthmonitor"`
}

func (o UpdateHealthmonitorRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateHealthmonitorRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateHealthmonitorRequestBody", string(data)}, " ")
}
