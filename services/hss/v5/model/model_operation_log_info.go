package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OperationLogInfo struct {

	// 主机ID
	HostId *string `json:"host_id,omitempty"`

	// 主机名称
	HostName *string `json:"host_name,omitempty"`

	// 备份名称
	BackupName *string `json:"backup_name,omitempty"`

	// 恢复进度（百分比）
	Process *int32 `json:"process,omitempty"`

	// 恢复状态，包含如下：   - success : 成功   - skipped : 跳过   - failed : 失败   - running : 正在运行   - timeout : 超时   - waiting : 等待
	Status *string `json:"status,omitempty"`

	// 任务开始时间
	StartedAt *string `json:"started_at,omitempty"`

	// 任务结束时间
	EndedAt *string `json:"ended_at,omitempty"`

	ErrorInfo *ErrorInfo `json:"error_info,omitempty"`
}

func (o OperationLogInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperationLogInfo struct{}"
	}

	return strings.Join([]string{"OperationLogInfo", string(data)}, " ")
}
