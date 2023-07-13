package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ItemResultDetailVo struct {

	// 状态
	Status *int32 `json:"status,omitempty"`

	// 检查项风险等级
	Level *int32 `json:"level,omitempty"`

	// 检查项修复建议
	Suggestion *string `json:"suggestion,omitempty"`

	// 检查项response
	Response *string `json:"response,omitempty"`

	// 检查项ID
	CheckId *string `json:"check_id,omitempty"`

	// 检查项名称
	CheckName *string `json:"check_name,omitempty"`

	// 检查项英文名称
	CheckNameEn *string `json:"check_name_en,omitempty"`

	// 检查项修复建议URL
	SuggestionUrl *string `json:"suggestion_url,omitempty"`
}

func (o ItemResultDetailVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ItemResultDetailVo struct{}"
	}

	return strings.Join([]string{"ItemResultDetailVo", string(data)}, " ")
}
