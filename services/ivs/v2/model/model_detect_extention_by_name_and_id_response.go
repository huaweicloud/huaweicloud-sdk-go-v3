package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DetectExtentionByNameAndIdResponse struct {
	Meta *Meta `json:"meta,omitempty"`

	Result         *IvsExtentionByNameAndIdResponseBodyResult `json:"result,omitempty"`
	HttpStatusCode int                                        `json:"-"`
}

func (o DetectExtentionByNameAndIdResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DetectExtentionByNameAndIdResponse struct{}"
	}

	return strings.Join([]string{"DetectExtentionByNameAndIdResponse", string(data)}, " ")
}
