package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConflictSectionLineDto struct {

	// 行
	LineCode *string `json:"line_code,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 旧行
	OldLine *int32 `json:"old_line,omitempty"`

	// 新行
	NewLine *int32 `json:"new_line,omitempty"`

	// 文本
	Text *string `json:"text,omitempty"`

	MetaData *ConflictSectionLineMetaDataDto `json:"meta_data,omitempty"`

	// 富文本
	RichText *string `json:"rich_text,omitempty"`

	// 可接受建议
	CanReceiveSuggestion *bool `json:"can_receive_suggestion,omitempty"`
}

func (o ConflictSectionLineDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConflictSectionLineDto struct{}"
	}

	return strings.Join([]string{"ConflictSectionLineDto", string(data)}, " ")
}
