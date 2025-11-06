package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MergeRequestCustomAverageEvaluationDto struct {

	// **参数解释：** 自定义评价维度id。
	EvaluationTypeId *int32 `json:"evaluation_type_id,omitempty"`

	// **参数解释：** 自定义评价维度名称。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 自定义评价维度平均分。
	Level *float64 `json:"level,omitempty"`
}

func (o MergeRequestCustomAverageEvaluationDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MergeRequestCustomAverageEvaluationDto struct{}"
	}

	return strings.Join([]string{"MergeRequestCustomAverageEvaluationDto", string(data)}, " ")
}
