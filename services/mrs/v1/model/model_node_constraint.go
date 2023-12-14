package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeConstraint 节点限制
type NodeConstraint struct {

	// 其他限制
	Other map[string]interface{} `json:"other,omitempty"`

	// 最少节点数
	MinNodeNum *int32 `json:"min_node_num,omitempty"`

	// 最多节点数
	MaxNodeNum *int32 `json:"max_node_num,omitempty"`

	// 最少核心数
	MinCoreNum map[string]int32 `json:"min_core_num,omitempty"`

	// 最小内存容量
	MinMemSize map[string]int32 `json:"min_mem_size,omitempty"`

	// 最小磁盘容量
	MinDiskSize *int32 `json:"min_disk_size,omitempty"`

	// 最大节点组数
	MaxNodeGroupNum *int32 `json:"max_node_group_num,omitempty"`

	// 最小数据卷容量
	MinDataVolumeTotalSize map[string]int32 `json:"min_data_volume_total_size,omitempty"`

	// 磁盘类型限制，包含当前节点组所支持的磁盘类型
	DiskTypeConstraint map[string]string `json:"disk_type_constraint,omitempty"`

	// 最小系统磁盘大小
	MinRootDiskSize *int32 `json:"min_root_disk_size,omitempty"`
}

func (o NodeConstraint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeConstraint struct{}"
	}

	return strings.Join([]string{"NodeConstraint", string(data)}, " ")
}
