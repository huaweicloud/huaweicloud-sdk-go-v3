package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StepConfig 作业配置。
type StepConfig struct {
	JobExecution *JobExecution `json:"job_execution"`
}

func (o StepConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StepConfig struct{}"
	}

	return strings.Join([]string{"StepConfig", string(data)}, " ")
}
