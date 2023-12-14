package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogicalClusterInfo This is a auto create Body Object
type LogicalClusterInfo struct {

	// 逻辑集群ID
	LogicalClusterId *string `json:"logical_cluster_id,omitempty"`

	// 逻辑集群名称
	LogicalClusterName *string `json:"logical_cluster_name,omitempty"`

	// 逻辑集群主机信息
	ClusterRings *[]ClusterRing `json:"cluster_rings,omitempty"`

	// 逻辑集群状态
	Status *string `json:"status,omitempty"`

	// 是否为第一个逻辑集群。第1个创建或者转换的逻辑集群不能删除，因为其中包含了一些系统视图
	FirstLogicalCluster *bool `json:"first_logical_cluster,omitempty"`

	ActionInfo *ActionInfo `json:"action_info,omitempty"`

	// 是否允许编辑
	EditEnable *bool `json:"edit_enable,omitempty"`

	// 是否允许重启
	RestartEnable *bool `json:"restart_enable,omitempty"`

	// 是否允许删除
	DeleteEnable *bool `json:"delete_enable,omitempty"`

	// 是否允许弹性伸缩
	AddToElastic *bool `json:"add_to_elastic,omitempty"`

	// 逻辑集群模式
	Mode *string `json:"mode,omitempty"`

	// 等待被销毁
	WaitingForKilling *int32 `json:"waiting_for_killing,omitempty"`

	// 集群类型
	ClusterType *string `json:"cluster_type,omitempty"`
}

func (o LogicalClusterInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogicalClusterInfo struct{}"
	}

	return strings.Join([]string{"LogicalClusterInfo", string(data)}, " ")
}
