package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 表达式
type Expression struct {

	// 公式，最多1024个字符(分析任务单输出场景，配合TransformModel或AggregateModel的output_property使用)
	Formula *string `json:"formula,omitempty"`

	// 带名称的公式
	Formulas *[]NamedFormula `json:"formulas,omitempty"`

	// 时间范围，调度时间往前的时间范围，比如1m表示调度时间往前1分钟到调度时间的时间范围，正则：\"1m|5m|15m|1h\"
	TimeRange string `json:"time_range"`
}

func (o Expression) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Expression struct{}"
	}

	return strings.Join([]string{"Expression", string(data)}, " ")
}
