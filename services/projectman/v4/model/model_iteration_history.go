package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IterationHistory struct {

	// 迭代ID
	IterationId *string `json:"iteration_id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	Operator *IterationHistoryOperator `json:"operator,omitempty"`

	// 操作类型
	Operate *string `json:"operate,omitempty"`

	// 操作时间
	OperateTime *string `json:"operate_time,omitempty"`

	// 操作详情
	Details *[]IterationHistoryDetails `json:"details,omitempty"`
}

func (o IterationHistory) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IterationHistory struct{}"
	}

	return strings.Join([]string{"IterationHistory", string(data)}, " ")
}
