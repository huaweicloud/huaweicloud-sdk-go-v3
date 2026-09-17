package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeviceIoTdbPushInfoDetail 创建设备数据推送IoTDB
type DeviceIoTdbPushInfoDetail struct {

	// 数据存储的存储组
	StorageGroup *string `json:"storage_group,omitempty"`

	// 数据格式转换类型
	Format *string `json:"format,omitempty"`
}

func (o DeviceIoTdbPushInfoDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeviceIoTdbPushInfoDetail struct{}"
	}

	return strings.Join([]string{"DeviceIoTdbPushInfoDetail", string(data)}, " ")
}
