package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DetectExtentionByIdCardImageResponse struct {
	Meta *Meta `json:"meta,omitempty"`

	Result         *IvsExtentionByIdCardImageResponseBodyResult `json:"result,omitempty"`
	HttpStatusCode int                                          `json:"-"`
}

func (o DetectExtentionByIdCardImageResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DetectExtentionByIdCardImageResponse struct{}"
	}

	return strings.Join([]string{"DetectExtentionByIdCardImageResponse", string(data)}, " ")
}
