package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAimMsgTemplateRequest Request Object
type UpdateAimMsgTemplateRequest struct {

	// 短信模板ID。
	TemplateId string `json:"template_id"`

	// 请求体参数类型，该字段必须设置为：application/json。
	ContentType string `json:"Content-Type"`

	Body *MsgTemplateRequest `json:"body,omitempty"`
}

func (o UpdateAimMsgTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAimMsgTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateAimMsgTemplateRequest", string(data)}, " ")
}
