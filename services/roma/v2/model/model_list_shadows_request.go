package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListShadowsRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 设备ID
	DeviceId int32 `json:"device_id"`
}

func (o ListShadowsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListShadowsRequest struct{}"
	}

	return strings.Join([]string{"ListShadowsRequest", string(data)}, " ")
}
