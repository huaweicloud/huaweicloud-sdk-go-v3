package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DetectStandardByIdCardImageRequest struct {
	Body *IvsStandardByIdCardImageRequestBody `json:"body,omitempty"`
}

func (o DetectStandardByIdCardImageRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DetectStandardByIdCardImageRequest struct{}"
	}

	return strings.Join([]string{"DetectStandardByIdCardImageRequest", string(data)}, " ")
}
