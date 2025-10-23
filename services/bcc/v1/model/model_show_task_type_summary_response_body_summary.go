package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskTypeSummaryResponseBodySummary 统计数据
type ShowTaskTypeSummaryResponseBodySummary struct {

	// 总任务数据
	TotalTaskCount int32 `json:"total_task_count"`

	// 备份任务总数
	BackupTaskCount int32 `json:"backup_task_count"`

	// 复制任务总数
	ReplicationTaskCount int32 `json:"replication_task_count"`

	// 恢复任务总数
	RestoreTaskCount int32 `json:"restore_task_count"`

	// 删除备份任务总数
	DeleteTaskCount int32 `json:"delete_task_count"`

	// 删除存储库任务总数
	VaultDeleteTaskCount int32 `json:"vault_delete_task_count"`

	// 移除存储库资源总数
	RemoveResourceTaskCount int32 `json:"remove_resource_task_count"`
}

func (o ShowTaskTypeSummaryResponseBodySummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskTypeSummaryResponseBodySummary struct{}"
	}

	return strings.Join([]string{"ShowTaskTypeSummaryResponseBodySummary", string(data)}, " ")
}
