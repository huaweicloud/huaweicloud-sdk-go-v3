package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 各项内容审核结果
type CategorySuggestions struct {

	// 政治人物审核
	Politics *string `json:"politics,omitempty" xml:"politics"`

	// 暴恐内容审核
	Terrorism *string `json:"terrorism,omitempty" xml:"terrorism"`

	// 情色内容审核
	Porn *string `json:"porn,omitempty" xml:"porn"`
}

func (o CategorySuggestions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CategorySuggestions struct{}"
	}

	return strings.Join([]string{"CategorySuggestions", string(data)}, " ")
}
