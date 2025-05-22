package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListFactoryPendingItemsRespData struct {

	// 已发布节点版本
	DeployedVersion *int32 `json:"deployed_version,omitempty"`

	// 任务名称
	ItemName *string `json:"item_name,omitempty"`

	// 打包状态，0表示未打包。
	PackageFlag *int32 `json:"package_flag,omitempty"`

	// 待发布包ID
	PendingItemId *string `json:"pending_item_id,omitempty"`

	// 待发布包版本
	PendingVersion *int32 `json:"pending_version,omitempty"`

	// 租户id+空间id
	ProjectId *string `json:"project_id,omitempty"`

	// 提交时间
	SubmitTimestamp *int64 `json:"submit_timestamp,omitempty"`

	// 提交人id
	SubmitUserId *string `json:"submit_user_id,omitempty"`

	// 提交人名称
	SubmitUserName *string `json:"submit_user_name,omitempty"`

	// 任务id
	TaskId *string `json:"task_id,omitempty"`

	// 任务类型，取值为1和2。 1: job 2: script
	TaskType *int32 `json:"task_type,omitempty"`

	// 变更类型，取值为1，2和3. 1: 新增 2: 变更 3: 删除
	UpdateType *int32 `json:"update_type,omitempty"`
}

func (o ListFactoryPendingItemsRespData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFactoryPendingItemsRespData struct{}"
	}

	return strings.Join([]string{"ListFactoryPendingItemsRespData", string(data)}, " ")
}
