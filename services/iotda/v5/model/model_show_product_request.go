package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowProductRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`

	ProductId string `json:"product_id"`

	AppId *string `json:"app_id,omitempty"`
}

func (o ShowProductRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowProductRequest struct{}"
	}

	return strings.Join([]string{"ShowProductRequest", string(data)}, " ")
}
