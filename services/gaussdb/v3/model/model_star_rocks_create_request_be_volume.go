package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StarRocksCreateRequestBeVolume BE节点存储规格。
type StarRocksCreateRequestBeVolume struct {

	// 磁盘类型。通过查询HTAP引擎资源返回消息获取。
	IoType string `json:"io_type"`

	// 磁盘容量，单位GB 增长的步长：10GB。
	CapacityInGb int32 `json:"capacity_in_gb"`
}

func (o StarRocksCreateRequestBeVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StarRocksCreateRequestBeVolume struct{}"
	}

	return strings.Join([]string{"StarRocksCreateRequestBeVolume", string(data)}, " ")
}
