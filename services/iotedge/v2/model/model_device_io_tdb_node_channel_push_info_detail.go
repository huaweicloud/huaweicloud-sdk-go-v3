package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeviceIoTdbNodeChannelPushInfoDetail 创建设备数据推送IoTDB
type DeviceIoTdbNodeChannelPushInfoDetail struct {

	// 数据存储的存储组
	StorageGroup *string `json:"storage_group,omitempty"`
}

func (o DeviceIoTdbNodeChannelPushInfoDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeviceIoTdbNodeChannelPushInfoDetail struct{}"
	}

	return strings.Join([]string{"DeviceIoTdbNodeChannelPushInfoDetail", string(data)}, " ")
}
