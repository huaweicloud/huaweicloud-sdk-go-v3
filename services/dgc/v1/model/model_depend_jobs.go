package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DependJobs struct {

	// 依赖的作业名称列表，必须依赖已存在的作业.
	Jobs []string `json:"jobs"`

	// 依赖周期
	DependPeriod *string `json:"dependPeriod,omitempty"`

	// 依赖作业任务执行失败处理策略
	DependFailPolicy *string `json:"dependFailPolicy,omitempty"`
}

func (o DependJobs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DependJobs struct{}"
	}

	return strings.Join([]string{"DependJobs", string(data)}, " ")
}
