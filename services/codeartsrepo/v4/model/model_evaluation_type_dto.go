package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EvaluationTypeDto 自定义评价返回体
type EvaluationTypeDto struct {

	// **参数解释：** 自定义评价主键id
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 自定义评价名称
	Name *string `json:"name,omitempty"`
}

func (o EvaluationTypeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EvaluationTypeDto struct{}"
	}

	return strings.Join([]string{"EvaluationTypeDto", string(data)}, " ")
}
