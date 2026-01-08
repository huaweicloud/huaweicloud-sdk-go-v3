package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Iops 云硬盘每秒进行读写的操作次数。
type Iops struct {

	// 产品ID。
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// 最小值。
	Min *int32 `json:"min,omitempty"`

	// 最大值。
	Max *int32 `json:"max,omitempty"`

	// 可取值范围，<=rangeBySize乘size。
	RangeBySize float32 `json:"range_by_size,omitempty"`
}

func (o Iops) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Iops struct{}"
	}

	return strings.Join([]string{"Iops", string(data)}, " ")
}
