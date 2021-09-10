package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DetectExtentionByIdCardImageRequest struct {
	Body *IvsExtentionByIdCardImageRequestBody `json:"body,omitempty"`
}

func (o DetectExtentionByIdCardImageRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DetectExtentionByIdCardImageRequest struct{}"
	}

	return strings.Join([]string{"DetectExtentionByIdCardImageRequest", string(data)}, " ")
}
