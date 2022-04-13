package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowProductTemplateRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 产品模板ID

	ProductTemplateId int32 `json:"product_template_id"`
}

func (o ShowProductTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProductTemplateRequest struct{}"
	}

	return strings.Join([]string{"ShowProductTemplateRequest", string(data)}, " ")
}
