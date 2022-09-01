package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SQL语句不支持转换的详情。
type UnSupportedItem struct {

	// SQL语句不支持转换的原因。
	Reason *string `json:"reason,omitempty" xml:"reason"`

	// SQL语句不支持转换的建议。
	Suggestion *string `json:"suggestion,omitempty" xml:"suggestion"`

	// 行号。
	LineNumber *int32 `json:"line_number,omitempty" xml:"line_number"`

	// 位置。
	Position *int32 `json:"position,omitempty" xml:"position"`
}

func (o UnSupportedItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnSupportedItem struct{}"
	}

	return strings.Join([]string{"UnSupportedItem", string(data)}, " ")
}
