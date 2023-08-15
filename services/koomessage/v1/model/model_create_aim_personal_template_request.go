package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAimPersonalTemplateRequest Request Object
type CreateAimPersonalTemplateRequest struct {

	// 请求体参数类型，该字段必须设置为：application/json。
	ContentType string `json:"Content-Type"`

	Body *CreateAimPersonalTemplateRequestBody `json:"body,omitempty"`
}

func (o CreateAimPersonalTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAimPersonalTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateAimPersonalTemplateRequest", string(data)}, " ")
}
