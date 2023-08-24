package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAimMsgTemplateRequest Request Object
type DeleteAimMsgTemplateRequest struct {

	// 短信模板ID。
	TemplateId string `json:"template_id"`

	// 请求体参数类型，该字段必须设置为：application/json。
	ContentType string `json:"Content-Type"`
}

func (o DeleteAimMsgTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAimMsgTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteAimMsgTemplateRequest", string(data)}, " ")
}
