package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListMessageTemplatesRequest struct {
	// 偏移量，偏移量为一个大于0小于资源总个数的整数，表示查询该偏移量后面的所有的资源，默认值为0。

	Offset *int32 `json:"offset,omitempty"`
	// 取值范围：1~100，取值一般为10，20，50。功能说明：每页返回的资源个数。默认值为100。

	Limit *int32 `json:"limit,omitempty"`
	// 模板的名称。  只能包含大写字母、小写字母、数字、-和_，且必须由大写字母、小写字母或数字开头，长度在1到64个字符之间。

	MessageTemplateName *string `json:"message_template_name,omitempty"`
	// 模板支持的协议类型。  目前支持的协议包括：  “default”：默认协议。  “email”：邮件传输协议。  “sms”：短信传输协议。  “functionstage”：FunctionGraph（函数）传输协议。  “dms”：DMS传输协议。  “http”、“https”：HTTP/HTTPS传输协议。

	Protocol *string `json:"protocol,omitempty"`
}

func (o ListMessageTemplatesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListMessageTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListMessageTemplatesRequest", string(data)}, " ")
}
