package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowPackageDetailRespTaskDetails struct {

	// 发布状态，1:待审批,2:成功,3:失败, 5:发布中
	DeployStatus *int32 `json:"deploy_status,omitempty"`

	// 已发布节点版本
	DeployedVersion *int32 `json:"deployed_version,omitempty"`

	// 发布任务名称
	ItemName *string `json:"item_name,omitempty"`

	// 发布任务ID
	PendingItemId *string `json:"pending_item_id,omitempty"`

	// 节点版本
	PendingVersion *int32 `json:"pending_version,omitempty"`

	// 具体脚本ID
	ScriptId *string `json:"script_id,omitempty"`

	// 作业ID
	TaskId *string `json:"task_id,omitempty"`

	// 作业启动状态，2：成功，3：失败
	StartJobStatus *int32 `json:"start_job_status,omitempty"`

	// 提交时间戳，13位时间戳
	SubmitTimestamp *int64 `json:"submit_timestamp,omitempty"`

	// 提交人id
	SubmitUserId *string `json:"submit_user_id,omitempty"`

	// 提交人名称
	SubmitUserName *string `json:"submit_user_name,omitempty"`

	// 任务类型（1-作业，2-脚本，3-资源）
	TaskType *int32 `json:"task_type,omitempty"`

	// 变更类型，默认值1，（1-新增，2-修改，3-删除）
	UpdateType *int32 `json:"update_type,omitempty"`
}

func (o ShowPackageDetailRespTaskDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPackageDetailRespTaskDetails struct{}"
	}

	return strings.Join([]string{"ShowPackageDetailRespTaskDetails", string(data)}, " ")
}
