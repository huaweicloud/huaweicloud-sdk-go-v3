package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFunctionTemplateRequest Request Object
type ShowFunctionTemplateRequest struct {

	// 指定模板id。
	TemplateId string `json:"template_id"`

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`
}

func (o ShowFunctionTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFunctionTemplateRequest struct{}"
	}

	return strings.Join([]string{"ShowFunctionTemplateRequest", string(data)}, " ")
}
