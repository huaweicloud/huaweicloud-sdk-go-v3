package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogicalClusterTaskInfo 逻辑集群任务信息
type LogicalClusterTaskInfo struct {

	// 任务类型
	Type *string `json:"type,omitempty"`

	// 逻辑集群名称
	LogicalClusterName *string `json:"logical_cluster_name,omitempty"`

	// 任务开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 任务结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 任务执行结果
	Result *string `json:"result,omitempty"`

	// 任务执行日志
	Log *string `json:"log,omitempty"`
}

func (o LogicalClusterTaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogicalClusterTaskInfo struct{}"
	}

	return strings.Join([]string{"LogicalClusterTaskInfo", string(data)}, " ")
}
