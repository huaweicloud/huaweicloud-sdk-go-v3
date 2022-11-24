package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 表达式
type Formula struct {

	// 公式，最多1024个字符(分析任务单输出场景，配合TransformModel或AggregateModel的output_property使用)
	Formula *string `json:"formula,omitempty"`

	// 带名称的公式
	Formulas *[]NamedFormula `json:"formulas,omitempty"`
}

func (o Formula) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Formula struct{}"
	}

	return strings.Join([]string{"Formula", string(data)}, " ")
}
