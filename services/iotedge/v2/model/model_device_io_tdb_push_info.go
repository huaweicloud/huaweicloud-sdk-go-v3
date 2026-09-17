package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeviceIoTdbPushInfo 创建设备数据推送IoTDB
type DeviceIoTdbPushInfo struct {

	// 数据存储的存储组
	StorageGroup string `json:"storage_group"`

	// 数据格式转换类型
	Format string `json:"format"`
}

func (o DeviceIoTdbPushInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeviceIoTdbPushInfo struct{}"
	}

	return strings.Join([]string{"DeviceIoTdbPushInfo", string(data)}, " ")
}
