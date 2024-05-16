package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecurityMemberSyncTaskRequest Request Object
type ShowSecurityMemberSyncTaskRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 用户同步任务id。
	TaskId string `json:"task_id"`
}

func (o ShowSecurityMemberSyncTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityMemberSyncTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowSecurityMemberSyncTaskRequest", string(data)}, " ")
}
