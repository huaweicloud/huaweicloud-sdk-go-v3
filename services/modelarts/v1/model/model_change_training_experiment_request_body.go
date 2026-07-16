package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeTrainingExperimentRequestBody 训练实验请求。
type ChangeTrainingExperimentRequestBody struct {

	// 训练实验名称。
	Name *string `json:"name,omitempty"`

	// 训练实验描述。
	Description *string `json:"description,omitempty"`
}

func (o ChangeTrainingExperimentRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeTrainingExperimentRequestBody struct{}"
	}

	return strings.Join([]string{"ChangeTrainingExperimentRequestBody", string(data)}, " ")
}
