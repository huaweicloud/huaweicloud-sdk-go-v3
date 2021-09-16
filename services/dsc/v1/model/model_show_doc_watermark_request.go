package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowDocWatermarkRequest struct {
	Body *ShowDocWatermarkRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o ShowDocWatermarkRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDocWatermarkRequest struct{}"
	}

	return strings.Join([]string{"ShowDocWatermarkRequest", string(data)}, " ")
}
