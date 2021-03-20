package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateProductRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`

	Body *AddProduct `json:"body,omitempty"`
}

func (o CreateProductRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateProductRequest struct{}"
	}

	return strings.Join([]string{"CreateProductRequest", string(data)}, " ")
}
