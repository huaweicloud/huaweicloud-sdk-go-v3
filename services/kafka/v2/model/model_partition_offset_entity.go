package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PartitionOffsetEntity 分区消费位点详情
type PartitionOffsetEntity struct {

	// 分区
	Partition *int32 `json:"partition,omitempty"`

	// 消费位点
	Offset *int32 `json:"offset,omitempty"`
}

func (o PartitionOffsetEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PartitionOffsetEntity struct{}"
	}

	return strings.Join([]string{"PartitionOffsetEntity", string(data)}, " ")
}
