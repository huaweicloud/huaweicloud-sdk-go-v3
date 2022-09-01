package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateMessageTemplateResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty" xml:"request_id"`

	// 模板唯一的资源标识。
	MessageTemplateId *string `json:"message_template_id,omitempty" xml:"message_template_id"`
	HttpStatusCode    int     `json:"-"`
}

func (o CreateMessageTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMessageTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreateMessageTemplateResponse", string(data)}, " ")
}
