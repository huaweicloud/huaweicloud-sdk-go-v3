package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListRecordCallbackConfigsRequest struct {
	// 直播推流域名

	PublishDomain *string `json:"publish_domain,omitempty"`
	// 流应用名称

	App *string `json:"app,omitempty"`
	// 分页编号，从0开始算

	Offset *int32 `json:"offset,omitempty"`
	// 每页记录数，取值范围[1,100]

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListRecordCallbackConfigsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListRecordCallbackConfigsRequest struct{}"
	}

	return strings.Join([]string{"ListRecordCallbackConfigsRequest", string(data)}, " ")
}
