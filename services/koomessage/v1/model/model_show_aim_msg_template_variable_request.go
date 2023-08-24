package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAimMsgTemplateVariableRequest Request Object
type ShowAimMsgTemplateVariableRequest struct {

	// 短信模板ID。
	TemplateId string `json:"template_id"`

	// 请求体参数类型，该字段必须设置为：application/json。
	ContentType string `json:"Content-Type"`
}

func (o ShowAimMsgTemplateVariableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAimMsgTemplateVariableRequest struct{}"
	}

	return strings.Join([]string{"ShowAimMsgTemplateVariableRequest", string(data)}, " ")
}
