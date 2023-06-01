package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListMessageTemplateDetailsRequest struct {

	// 模板唯一的资源标识，可通过查询[消息模板列表](ListMessageTemplates.xml)获取该标识。
	MessageTemplateId string `json:"message_template_id"`
}

func (o ListMessageTemplateDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMessageTemplateDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListMessageTemplateDetailsRequest", string(data)}, " ")
}
