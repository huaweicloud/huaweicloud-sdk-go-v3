package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteTranscodingsTemplateRequest struct {
	// 播放域名

	Domain string `json:"domain"`
	// 应用名称

	AppName string `json:"app_name"`
}

func (o DeleteTranscodingsTemplateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteTranscodingsTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteTranscodingsTemplateRequest", string(data)}, " ")
}
