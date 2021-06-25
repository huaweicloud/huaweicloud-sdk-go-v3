package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteWatermarkTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteWatermarkTemplateResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteWatermarkTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteWatermarkTemplateResponse", string(data)}, " ")
}
