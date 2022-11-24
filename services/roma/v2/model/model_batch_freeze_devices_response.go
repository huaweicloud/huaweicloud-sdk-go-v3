package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchFreezeDevicesResponse struct {

	// 下线成功设备列表
	Success *[]DeviceInfoSimple `json:"success,omitempty"`

	// 下线失败设备列表
	Failed         *[]DeviceInfoSimple `json:"failed,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o BatchFreezeDevicesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchFreezeDevicesResponse struct{}"
	}

	return strings.Join([]string{"BatchFreezeDevicesResponse", string(data)}, " ")
}
