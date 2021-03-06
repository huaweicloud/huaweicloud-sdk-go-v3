package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListMessageTemplateDetailsRequest struct {
	// 模板唯一的资源标识，可通过查询[消息模板列表](https://support.huaweicloud.com/api-smn/smn_api_53004.html)获取该标识。

	MessageTemplateId string `json:"message_template_id"`
}

func (o ListMessageTemplateDetailsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListMessageTemplateDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListMessageTemplateDetailsRequest", string(data)}, " ")
}
