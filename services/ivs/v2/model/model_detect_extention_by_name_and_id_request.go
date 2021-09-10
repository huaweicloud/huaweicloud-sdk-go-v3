package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DetectExtentionByNameAndIdRequest struct {
	Body *IvsExtentionByNameAndIdRequestBody `json:"body,omitempty"`
}

func (o DetectExtentionByNameAndIdRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DetectExtentionByNameAndIdRequest struct{}"
	}

	return strings.Join([]string{"DetectExtentionByNameAndIdRequest", string(data)}, " ")
}
