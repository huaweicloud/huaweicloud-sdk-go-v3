package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteWatermarkTemplateRequest struct {
	// 水印模板配置id

	Id string `json:"id"`
}

func (o DeleteWatermarkTemplateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteWatermarkTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteWatermarkTemplateRequest", string(data)}, " ")
}
