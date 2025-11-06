package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomEvaluationDto struct {

	// **参数解释：** 自定义评价id。
	Id *int32 `json:"id,omitempty"`

	// 评价类型id
	EvaluationTypeId *int32 `json:"evaluation_type_id,omitempty"`

	// 评价等级
	Level *int32 `json:"level,omitempty"`

	// 评价名称
	Name *string `json:"name,omitempty"`
}

func (o CustomEvaluationDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomEvaluationDto struct{}"
	}

	return strings.Join([]string{"CustomEvaluationDto", string(data)}, " ")
}
