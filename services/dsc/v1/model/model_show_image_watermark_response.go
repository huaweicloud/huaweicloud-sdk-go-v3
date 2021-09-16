package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowImageWatermarkResponse struct {
	// 暗水印内容，长度不超过32个字节

	Watermark      *string `json:"watermark,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowImageWatermarkResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowImageWatermarkResponse struct{}"
	}

	return strings.Join([]string{"ShowImageWatermarkResponse", string(data)}, " ")
}
