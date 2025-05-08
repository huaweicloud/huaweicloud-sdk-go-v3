package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogicalClusterPlanBo 用于提交逻辑集群定时增删计划
type LogicalClusterPlanBo struct {

	// 逻辑集群名字
	LogicalClusterName *string `json:"logical_cluster_name,omitempty"`

	// 逻辑集群绑定的用户，若绑定了主逻辑集群，不能绑定用户
	User *string `json:"user,omitempty"`

	// 逻辑集群节点的个数
	NodeNum *int32 `json:"node_num,omitempty"`

	// 逻辑集群的绑定的主逻辑集群，若绑定了用户，不能绑定主逻辑集群
	MainLogicalCluster *string `json:"main_logical_cluster,omitempty"`

	// 计划类型，取值范围为(once|periodicity)
	PlanType string `json:"plan_type"`

	// 逻辑集群定时增删计划起始时间
	StartTime *string `json:"start_time,omitempty"`

	// 逻辑集群定时增删计划终止时间
	EndTime *string `json:"end_time,omitempty"`

	// 逻辑集群定时增删计划细节
	Actions []LogicalClusterPlanActionsParam `json:"actions"`
}

func (o LogicalClusterPlanBo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogicalClusterPlanBo struct{}"
	}

	return strings.Join([]string{"LogicalClusterPlanBo", string(data)}, " ")
}
