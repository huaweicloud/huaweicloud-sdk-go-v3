package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Result struct {

	// 结果返回信息
	Message *interface{} `json:"message,omitempty"`

	// 元数据信息
	Schema *interface{} `json:"schema,omitempty"`

	// 每条结果的信息
	Rows *interface{} `json:"rows,omitempty"`

	// 结果行数
	RowCount *int64 `json:"rowCount,omitempty"`

	// 输入结果的行数。（dli等脚本执行会执行此结果）
	InputRowCount *int64 `json:"inputRowCount,omitempty"`

	// 结果行数。（dli等脚本执行会执行此结果）
	ResultCount *int64 `json:"resultCount,omitempty"`

	// 脚本运行时间
	Duration *float32 `json:"duration,omitempty"`

	// 脚本结果信息
	RawResult *interface{} `json:"rawResult,omitempty"`
}

func (o Result) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Result struct{}"
	}

	return strings.Join([]string{"Result", string(data)}, " ")
}
