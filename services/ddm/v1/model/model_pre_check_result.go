package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PreCheckResult struct {

	// 检查项名称。
	Name *string `json:"name,omitempty"`

	// 检查项结果。
	Status *string `json:"status,omitempty"`

	// 检查项提示信息。
	Note *string `json:"note,omitempty"`

	// 处理建议。
	HandlingSuggestion *string `json:"handling_suggestion,omitempty"`

	// 错误信息。
	ErrorMessage *string `json:"error_message,omitempty"`

	// 详细错误信息。
	ErrorDetailMessage *string `json:"error_detail_message,omitempty"`
}

func (o PreCheckResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PreCheckResult struct{}"
	}

	return strings.Join([]string{"PreCheckResult", string(data)}, " ")
}
