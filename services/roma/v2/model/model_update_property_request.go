package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdatePropertyRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 服务ID

	ServiceId string `json:"service_id"`
	// 属性/请求属性/响应属性ID

	PropertyId int32 `json:"property_id"`

	Body *UpdatePropertyRequestBody `json:"body,omitempty"`
}

func (o UpdatePropertyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePropertyRequest struct{}"
	}

	return strings.Join([]string{"UpdatePropertyRequest", string(data)}, " ")
}
