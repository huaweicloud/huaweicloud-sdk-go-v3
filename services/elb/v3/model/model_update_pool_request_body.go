package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type UpdatePoolRequestBody struct {
	Pool *UpdatePoolOption `json:"pool"`
}

func (o UpdatePoolRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdatePoolRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePoolRequestBody", string(data)}, " ")
}
