package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DefectEvents struct {

	// 调用链信息
	Events *[]DefectEvents `json:"events,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// fix suggestions
	FixSuggestions *[]string `json:"fix_suggestions,omitempty"`

	// 缺陷所在行
	Line *int32 `json:"line,omitempty"`

	// 缺陷结束行
	EndLine *int32 `json:"end_line,omitempty"`

	// 缺陷开始行
	CodeContextStartLine *int32 `json:"code_context_start_line,omitempty"`

	// 代码片段
	Main *bool `json:"main,omitempty"`

	// file path
	Path *string `json:"path,omitempty"`

	// 标签
	Tag *string `json:"tag,omitempty"`

	// main buggy code
	MainBuggyCode *string `json:"main_buggy_code,omitempty"`

	// code context
	CodeContext *string `json:"code_context,omitempty"`
}

func (o DefectEvents) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DefectEvents struct{}"
	}

	return strings.Join([]string{"DefectEvents", string(data)}, " ")
}
