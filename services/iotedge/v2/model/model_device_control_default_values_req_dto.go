package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeviceControlDefaultValuesReqDto 设备控制默认值请求结构体
type DeviceControlDefaultValuesReqDto struct {

	// 设备默认值数组
	Devices []DeviceDefaultValues `json:"devices"`
}

func (o DeviceControlDefaultValuesReqDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeviceControlDefaultValuesReqDto struct{}"
	}

	return strings.Join([]string{"DeviceControlDefaultValuesReqDto", string(data)}, " ")
}
