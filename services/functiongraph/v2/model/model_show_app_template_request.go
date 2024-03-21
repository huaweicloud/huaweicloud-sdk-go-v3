package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAppTemplateRequest Request Object
type ShowAppTemplateRequest struct {

	// 应用模板ID。
	Id string `json:"id"`

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`
}

func (o ShowAppTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppTemplateRequest struct{}"
	}

	return strings.Join([]string{"ShowAppTemplateRequest", string(data)}, " ")
}
