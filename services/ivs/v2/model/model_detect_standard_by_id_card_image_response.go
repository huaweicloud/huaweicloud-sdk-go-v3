package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DetectStandardByIdCardImageResponse struct {
	Meta *Meta `json:"meta,omitempty"`

	Result         *IvsStandardByIdCardImageResponseBodyResult `json:"result,omitempty"`
	HttpStatusCode int                                         `json:"-"`
}

func (o DetectStandardByIdCardImageResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DetectStandardByIdCardImageResponse struct{}"
	}

	return strings.Join([]string{"DetectStandardByIdCardImageResponse", string(data)}, " ")
}
