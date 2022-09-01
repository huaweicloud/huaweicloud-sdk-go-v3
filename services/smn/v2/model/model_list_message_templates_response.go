package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListMessageTemplatesResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty" xml:"request_id"`

	// 返回的模板个数。
	MessageTemplateCount *int32 `json:"message_template_count,omitempty" xml:"message_template_count"`

	// Message_template结构体数组。
	MessageTemplates *[]MessageTemplate `json:"message_templates,omitempty" xml:"message_templates"`
	HttpStatusCode   int                `json:"-"`
}

func (o ListMessageTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMessageTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListMessageTemplatesResponse", string(data)}, " ")
}
