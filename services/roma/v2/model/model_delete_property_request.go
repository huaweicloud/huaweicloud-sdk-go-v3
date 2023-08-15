package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePropertyRequest Request Object
type DeletePropertyRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 服务ID
	ServiceId string `json:"service_id"`

	// 属性/请求属性/响应属性ID
	PropertyId int32 `json:"property_id"`
}

func (o DeletePropertyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePropertyRequest struct{}"
	}

	return strings.Join([]string{"DeletePropertyRequest", string(data)}, " ")
}
