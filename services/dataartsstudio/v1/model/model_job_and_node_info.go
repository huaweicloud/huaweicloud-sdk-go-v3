package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobAndNodeInfo 作业算子基本信息
type JobAndNodeInfo struct {

	// 作业算子id
	TaskId *string `json:"task_id,omitempty"`

	// 作业算子名称
	JobName *string `json:"job_name,omitempty"`

	// 作业算子所在空间id
	WorkspaceId *string `json:"workspace_id,omitempty"`
}

func (o JobAndNodeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobAndNodeInfo struct{}"
	}

	return strings.Join([]string{"JobAndNodeInfo", string(data)}, " ")
}
