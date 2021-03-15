package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateProductRequest struct {
	InstanceId *string        `json:"Instance-Id,omitempty"`
	ProductId  string         `json:"product_id"`
	Body       *UpdateProduct `json:"body,omitempty"`
}

func (o UpdateProductRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateProductRequest struct{}"
	}

	return strings.Join([]string{"UpdateProductRequest", string(data)}, " ")
}
