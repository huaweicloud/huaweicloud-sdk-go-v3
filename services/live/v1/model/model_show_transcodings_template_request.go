package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowTranscodingsTemplateRequest struct {
	// 播放域名

	Domain string `json:"domain"`
	// 应用名称

	AppName *string `json:"app_name,omitempty"`
	// 分页编号，默认为0。

	Page *int32 `json:"page,omitempty"`
	// 每页记录数。  取值范围：1-100。  默认为10。

	Size *int32 `json:"size,omitempty"`
}

func (o ShowTranscodingsTemplateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowTranscodingsTemplateRequest struct{}"
	}

	return strings.Join([]string{"ShowTranscodingsTemplateRequest", string(data)}, " ")
}
