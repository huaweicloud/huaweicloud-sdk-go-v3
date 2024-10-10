package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobInformation struct {

	// 下游作业的工作空间名称
	WorkSpace *string `json:"workSpace,omitempty"`

	// 下游作业名称
	JobName *string `json:"jobName,omitempty"`

	// 下游作业ID
	JobId *int64 `json:"jobId,omitempty"`
}

func (o JobInformation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobInformation struct{}"
	}

	return strings.Join([]string{"JobInformation", string(data)}, " ")
}
