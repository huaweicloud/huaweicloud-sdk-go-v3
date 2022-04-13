package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateProductTemplateRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`

	Body *CreateProductTemplateRequestBody `json:"body,omitempty"`
}

func (o CreateProductTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProductTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateProductTemplateRequest", string(data)}, " ")
}
