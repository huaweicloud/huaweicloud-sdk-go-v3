package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchFreezeDevicesRequestBody struct {

	// 设备ID列表，自动向下取整
	DeviceIds []int32 `json:"device_ids"`
}

func (o BatchFreezeDevicesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchFreezeDevicesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchFreezeDevicesRequestBody", string(data)}, " ")
}
