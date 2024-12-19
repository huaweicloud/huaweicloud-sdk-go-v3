package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RetryFactoryJobInstanceBodyJobs struct {

	// 重跑的作业名称。
	JobName string `json:"job_name"`

	// 工作空间id
	WorkspaceId string `json:"workspace_id"`
}

func (o RetryFactoryJobInstanceBodyJobs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryFactoryJobInstanceBodyJobs struct{}"
	}

	return strings.Join([]string{"RetryFactoryJobInstanceBodyJobs", string(data)}, " ")
}
