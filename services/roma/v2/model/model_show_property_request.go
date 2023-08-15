package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPropertyRequest Request Object
type ShowPropertyRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 服务ID
	ServiceId string `json:"service_id"`

	// 属性/请求属性/响应属性ID
	PropertyId int32 `json:"property_id"`
}

func (o ShowPropertyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPropertyRequest struct{}"
	}

	return strings.Join([]string{"ShowPropertyRequest", string(data)}, " ")
}
