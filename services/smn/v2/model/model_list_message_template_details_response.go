package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMessageTemplateDetailsResponse Response Object
type ListMessageTemplateDetailsResponse struct {

	// 模板ID。
	MessageTemplateId *string `json:"message_template_id,omitempty"`

	// 模板名称。
	MessageTemplateName *string `json:"message_template_name,omitempty"`

	// 模板支持的协议类型。  目前支持的协议包括：  “default”：默认协议。  “email”：邮件传输协议。  “sms”：短信传输协议。  “functionstage”：FunctionGraph（函数）传输协议。  “http”、“https”：HTTP/HTTPS传输协议。
	Protocol *string `json:"protocol,omitempty"`

	// 模板tag列表。  是消息模板“{}”内的字段，在具体使用消息模板时，可根据实际情况替为该字段赋值。
	TagNames *[]string `json:"tag_names,omitempty"`

	// 模板创建时间。 时间格式为UTC时间，YYYY-MM-DDTHH:MM:SSZ。
	CreateTime *string `json:"create_time,omitempty"`

	// 模板最后更新时间。时间格式为UTC时间，YYYY-MM-DDTHH:MM:SSZ。
	UpdateTime *string `json:"update_time,omitempty"`

	// 模板内容。
	Content *string `json:"content,omitempty"`

	// 请求的唯一标识ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListMessageTemplateDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMessageTemplateDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListMessageTemplateDetailsResponse", string(data)}, " ")
}
