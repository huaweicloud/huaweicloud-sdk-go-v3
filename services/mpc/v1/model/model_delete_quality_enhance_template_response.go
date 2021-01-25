package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteQualityEnhanceTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteQualityEnhanceTemplateResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteQualityEnhanceTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteQualityEnhanceTemplateResponse", string(data)}, " ")
}
