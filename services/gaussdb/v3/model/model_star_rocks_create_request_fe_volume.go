package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StarRocksCreateRequestFeVolume FE节点存储规格。
type StarRocksCreateRequestFeVolume struct {

	// 磁盘类型。通过查询HTAP引擎资源返回消息获取。
	IoType string `json:"io_type"`

	// 磁盘容量，单位GB 增长的步长：10GB。
	CapacityInGb int32 `json:"capacity_in_gb"`
}

func (o StarRocksCreateRequestFeVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StarRocksCreateRequestFeVolume struct{}"
	}

	return strings.Join([]string{"StarRocksCreateRequestFeVolume", string(data)}, " ")
}
