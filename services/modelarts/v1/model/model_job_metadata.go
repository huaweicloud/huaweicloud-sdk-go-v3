package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobMetadata 训练作业元信息。
type JobMetadata struct {

	// 训练作业名称。限制为1-64位只含数字、字母、下划线和中划线的名称。
	Name string `json:"name"`

	// 指定作业所处的工作空间，默认值为“0”。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 对训练作业的描述，默认为“NULL”，字符串的长度限制为[0, 256]。
	Description *string `json:"description,omitempty"`

	// 训练作业高级功能配置，可选取值如下： - \"job_template\": \"Template RL\"（异构作业）。 - \"fault-tolerance/job-retry-num\": \"3\"（故障自动重启次数）。 - \"jupyter-lab/enable\": \"true\"（JupyterLab训练应用程序）
	Annotations map[string]string `json:"annotations,omitempty"`

	TrainingExperimentReference *TrainingExperimentRequest `json:"training_experiment_reference,omitempty"`
}

func (o JobMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobMetadata struct{}"
	}

	return strings.Join([]string{"JobMetadata", string(data)}, " ")
}
