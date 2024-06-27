package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClickHouseNodeInfoVolume 实例节点存储信息。
type ClickHouseNodeInfoVolume struct {

	// 实例节点存储类型。 取值范围： - SSD：超高IO - ESSD：极速型SSD
	Type string `json:"type"`

	// 实例节点存储大小。
	Size string `json:"size"`

	// 实例节点存储使用大小。
	Used *string `json:"used,omitempty"`

	// 实例节点存储IOPS大小。
	Iops *int32 `json:"iops,omitempty"`
}

func (o ClickHouseNodeInfoVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClickHouseNodeInfoVolume struct{}"
	}

	return strings.Join([]string{"ClickHouseNodeInfoVolume", string(data)}, " ")
}
