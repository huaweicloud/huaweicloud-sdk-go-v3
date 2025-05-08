package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogicalClusterPlanVo 逻辑增删集群计划
type LogicalClusterPlanVo struct {

	// 增删逻辑集群计划ID
	Id *string `json:"id,omitempty"`

	// 逻辑集群名字
	LogicalClusterName *string `json:"logical_cluster_name,omitempty"`

	// 逻辑集群节点个数
	NodeNum *int32 `json:"node_num,omitempty"`

	// 逻辑集群增删计划类型，取值范围为 (once|periodicity)，表示单次或周期性
	PlanType *string `json:"plan_type,omitempty"`

	// 逻辑集群增删计划状态，取值范围为(running|waiting|deleted|finished|disabled|failed)
	Status *string `json:"status,omitempty"`

	// 逻辑集群增删计划开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 逻辑集群增删计划结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 逻辑集群增删计划更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 逻辑集群增删计划绑定的用户
	User *string `json:"user,omitempty"`

	// 逻辑集群增删计划行动
	Actions *[]LogicalClusterPlanActions `json:"actions,omitempty"`
}

func (o LogicalClusterPlanVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogicalClusterPlanVo struct{}"
	}

	return strings.Join([]string{"LogicalClusterPlanVo", string(data)}, " ")
}
