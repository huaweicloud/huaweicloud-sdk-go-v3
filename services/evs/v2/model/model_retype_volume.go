package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetypeVolume 变更云硬盘类型
type RetypeVolume struct {

	// 变更至指定的云硬盘类型
	NewType string `json:"new_type"`

	// 云硬盘iops大小。
	Iops *int32 `json:"iops,omitempty"`

	// 云硬盘的吞吐量大小。
	Throughput *int32 `json:"throughput,omitempty"`
}

func (o RetypeVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetypeVolume struct{}"
	}

	return strings.Join([]string{"RetypeVolume", string(data)}, " ")
}
