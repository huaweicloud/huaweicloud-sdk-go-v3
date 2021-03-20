package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteProductRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`

	ProductId string `json:"product_id"`

	AppId *string `json:"app_id,omitempty"`
}

func (o DeleteProductRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteProductRequest struct{}"
	}

	return strings.Join([]string{"DeleteProductRequest", string(data)}, " ")
}
