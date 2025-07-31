package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssessProperty 评估指标详情
type AssessProperty struct {

	// 预期信息
	Expected *string `json:"expected,omitempty"`

	// 实际信息
	Actual *string `json:"actual,omitempty"`

	// 中文建议
	Suggestion *string `json:"suggestion,omitempty"`

	// 英文建议
	SuggestionEn *string `json:"suggestion_en,omitempty"`

	// 是否超出范围
	Status *int32 `json:"status,omitempty"`
}

func (o AssessProperty) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssessProperty struct{}"
	}

	return strings.Join([]string{"AssessProperty", string(data)}, " ")
}
