package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSnapshotPolicyRequestBody 更新快照策略请求体
type UpdateSnapshotPolicyRequestBody struct {

	// 服务器id
	HostId string `json:"host_id"`

	// 是否开启定时快照 |-closed：关闭， -opened：开启
	ScheduledTaskStatus string `json:"scheduled_task_status"`

	// 主机每天定时快照的时间，支持整点，0-23, 默认时间为零点。
	ScheduledTaskTime *int32 `json:"scheduled_task_time,omitempty"`
}

func (o UpdateSnapshotPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSnapshotPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSnapshotPolicyRequestBody", string(data)}, " ")
}
