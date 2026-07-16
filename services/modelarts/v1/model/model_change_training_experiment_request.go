package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeTrainingExperimentRequest Request Object
type ChangeTrainingExperimentRequest struct {

	// 训练实验ID。在训练作业创建时获取实验ID。
	ExperimentId string `json:"experiment_id"`

	Body *ChangeTrainingExperimentRequestBody `json:"body,omitempty"`
}

func (o ChangeTrainingExperimentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeTrainingExperimentRequest struct{}"
	}

	return strings.Join([]string{"ChangeTrainingExperimentRequest", string(data)}, " ")
}
