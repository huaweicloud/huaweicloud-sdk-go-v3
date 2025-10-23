package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowTaskTypeSummaryResponseBodyDatapoints struct {

	// 统计日期
	StatisticsDate *string `json:"statistics_date,omitempty"`

	// 当前统计周期内的所有类型任务数量
	TotalTaskCount *int32 `json:"total_task_count,omitempty"`

	// 当前统计周期内的备份任务数
	BackupTaskCount *int32 `json:"backup_task_count,omitempty"`

	// 当前统计周期内的复制任务数
	ReplicationTaskCount *int32 `json:"replication_task_count,omitempty"`

	// 当前统计周期内的恢复任务数
	RestoreTaskCount *int32 `json:"restore_task_count,omitempty"`

	// 当前统计周期内的删除备份任务数
	DeleteTaskCount *int32 `json:"delete_task_count,omitempty"`

	// 当前统计周期内的删除存储库任务数
	VaultDeleteTaskCount *int32 `json:"vault_delete_task_count,omitempty"`

	// 当前统计周期内的移除存储库资源数
	RemoveResourceTaskCount *int32 `json:"remove_resource_task_count,omitempty"`
}

func (o ShowTaskTypeSummaryResponseBodyDatapoints) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskTypeSummaryResponseBodyDatapoints struct{}"
	}

	return strings.Join([]string{"ShowTaskTypeSummaryResponseBodyDatapoints", string(data)}, " ")
}
