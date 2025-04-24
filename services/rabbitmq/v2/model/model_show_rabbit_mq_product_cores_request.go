package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRabbitMqProductCoresRequest Request Object
type ShowRabbitMqProductCoresRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 产品ID。
	ProductId string `json:"product_id"`
}

func (o ShowRabbitMqProductCoresRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRabbitMqProductCoresRequest struct{}"
	}

	return strings.Join([]string{"ShowRabbitMqProductCoresRequest", string(data)}, " ")
}
