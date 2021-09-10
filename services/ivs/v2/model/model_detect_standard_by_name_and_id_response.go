package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DetectStandardByNameAndIdResponse struct {
	Meta *Meta `json:"meta,omitempty"`

	Result         *IvsStandardByNameAndIdResponseBodyResult `json:"result,omitempty"`
	HttpStatusCode int                                       `json:"-"`
}

func (o DetectStandardByNameAndIdResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DetectStandardByNameAndIdResponse struct{}"
	}

	return strings.Join([]string{"DetectStandardByNameAndIdResponse", string(data)}, " ")
}
