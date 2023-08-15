package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeviceGroupTreeRequest Request Object
type ShowDeviceGroupTreeRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 应用ID
	AppId string `json:"app_id"`
}

func (o ShowDeviceGroupTreeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeviceGroupTreeRequest struct{}"
	}

	return strings.Join([]string{"ShowDeviceGroupTreeRequest", string(data)}, " ")
}
