package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateProductTopicRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 产品ID

	ProductId int32 `json:"product_id"`

	Body *CreateProductTopicRequestBody `json:"body,omitempty"`
}

func (o CreateProductTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProductTopicRequest struct{}"
	}

	return strings.Join([]string{"CreateProductTopicRequest", string(data)}, " ")
}
