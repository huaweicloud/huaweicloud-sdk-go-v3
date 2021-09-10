package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DetectStandardByNameAndIdRequest struct {
	Body *IvsStandardByNameAndIdRequestBody `json:"body,omitempty"`
}

func (o DetectStandardByNameAndIdRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DetectStandardByNameAndIdRequest struct{}"
	}

	return strings.Join([]string{"DetectStandardByNameAndIdRequest", string(data)}, " ")
}
