package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateProductTemplateRequest Request Object
type UpdateProductTemplateRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 产品模板ID
	ProductTemplateId int32 `json:"product_template_id"`

	Body *UpdateProductTemplateRequestBody `json:"body,omitempty"`
}

func (o UpdateProductTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProductTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateProductTemplateRequest", string(data)}, " ")
}
