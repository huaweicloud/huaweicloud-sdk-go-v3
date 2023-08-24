package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAimMsgTemplateRequest Request Object
type CreateAimMsgTemplateRequest struct {

	// 请求体参数类型，该字段必须设置为：application/json。
	ContentType string `json:"Content-Type"`

	Body *MsgTemplateRequest `json:"body,omitempty"`
}

func (o CreateAimMsgTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAimMsgTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateAimMsgTemplateRequest", string(data)}, " ")
}
