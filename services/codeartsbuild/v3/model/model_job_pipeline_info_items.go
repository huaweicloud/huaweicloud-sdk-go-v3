package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobPipelineInfoItems struct {

	// 构建执行SCM
	Scms *[]CreateBuildJobScm `json:"scms,omitempty"`

	// 参数
	Parameters *[]JobPipelineInfoParameters `json:"parameters,omitempty"`

	// 任务名称
	JobName *string `json:"job_name,omitempty"`

	// 任务名称信息
	JobNameMassage *string `json:"job_name_massage,omitempty"`

	// 任务名称正则
	JobNameRegex *string `json:"job_name_regex,omitempty"`

	// 任务名称正则
	SourceCode *string `json:"source_code,omitempty"`
}

func (o JobPipelineInfoItems) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobPipelineInfoItems struct{}"
	}

	return strings.Join([]string{"JobPipelineInfoItems", string(data)}, " ")
}
