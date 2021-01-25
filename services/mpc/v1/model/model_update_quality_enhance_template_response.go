package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateQualityEnhanceTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateQualityEnhanceTemplateResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateQualityEnhanceTemplateResponse struct{}"
	}

	return strings.Join([]string{"UpdateQualityEnhanceTemplateResponse", string(data)}, " ")
}
