package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateVmsTemplateRequest struct {

	// 请求体参数类型，该字段必须设置为：application/json。
	ContentType string `json:"Content-Type"`

	Body *CreateVmsTemplateRequestBody `json:"body,omitempty"`
}

func (o CreateVmsTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVmsTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateVmsTemplateRequest", string(data)}, " ")
}
