package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Throughput 云硬盘每秒成功传送的数据量，即读取和写入的数据量。
type Throughput struct {

	// 产品ID。
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// 最小值。
	Min *int32 `json:"min,omitempty"`

	// 最大值。
	Max *int32 `json:"max,omitempty"`

	// 可取值范围，<=rangeByIops乘iops。
	RangeByIops float32 `json:"range_by_iops,omitempty"`
}

func (o Throughput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Throughput struct{}"
	}

	return strings.Join([]string{"Throughput", string(data)}, " ")
}
