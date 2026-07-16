package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeTrainingExperimentResponse Response Object
type ChangeTrainingExperimentResponse struct {

	// 训练实验名称。
	Name *string `json:"name,omitempty"`

	// 训练实验描述。
	Description    *string `json:"description,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeTrainingExperimentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeTrainingExperimentResponse struct{}"
	}

	return strings.Join([]string{"ChangeTrainingExperimentResponse", string(data)}, " ")
}
