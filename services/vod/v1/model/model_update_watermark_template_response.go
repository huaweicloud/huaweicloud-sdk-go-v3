package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateWatermarkTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateWatermarkTemplateResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateWatermarkTemplateResponse struct{}"
	}

	return strings.Join([]string{"UpdateWatermarkTemplateResponse", string(data)}, " ")
}
