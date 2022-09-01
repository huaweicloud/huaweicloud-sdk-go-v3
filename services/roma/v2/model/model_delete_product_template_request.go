package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteProductTemplateRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 产品模板ID
	ProductTemplateId int32 `json:"product_template_id" xml:"product_template_id"`
}

func (o DeleteProductTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProductTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteProductTemplateRequest", string(data)}, " ")
}
