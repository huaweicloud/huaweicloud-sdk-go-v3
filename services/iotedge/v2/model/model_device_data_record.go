package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeviceDataRecord 设备上报数据打印开关
type DeviceDataRecord struct {

	// 磁盘配额，单位MB，参考值，只能保证在这个值左右
	DiskQuota int32 `json:"disk_quota"`

	// 老化时间，日志压缩文件名时间戳老于这个时间就会发生老化删除
	Age int32 `json:"age"`

	// 配置开关，true启用数据打印，false不启用数据打印
	State string `json:"state"`
}

func (o DeviceDataRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeviceDataRecord struct{}"
	}

	return strings.Join([]string{"DeviceDataRecord", string(data)}, " ")
}
