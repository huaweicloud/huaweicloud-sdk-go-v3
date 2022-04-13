package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListDevicesInProductRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 产品ID

	ProductId int32 `json:"product_id"`
}

func (o ListDevicesInProductRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDevicesInProductRequest struct{}"
	}

	return strings.Join([]string{"ListDevicesInProductRequest", string(data)}, " ")
}
