package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowImageWatermarkRequest struct {
	Body *ShowImageWatermarkRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o ShowImageWatermarkRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowImageWatermarkRequest struct{}"
	}

	return strings.Join([]string{"ShowImageWatermarkRequest", string(data)}, " ")
}
