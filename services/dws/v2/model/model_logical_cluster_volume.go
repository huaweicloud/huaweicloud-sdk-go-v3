package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogicalClusterVolume 逻辑集群磁盘信息
type LogicalClusterVolume struct {

	// 逻辑集群名称
	LogicalClusterName *string `json:"logical_cluster_name,omitempty"`

	// 磁盘使用量
	Usage *string `json:"usage,omitempty"`

	// 磁盘总量
	Total *string `json:"total,omitempty"`

	// 磁盘使用比例
	Percent *string `json:"percent,omitempty"`
}

func (o LogicalClusterVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogicalClusterVolume struct{}"
	}

	return strings.Join([]string{"LogicalClusterVolume", string(data)}, " ")
}
