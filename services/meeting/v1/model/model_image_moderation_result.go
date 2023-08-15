package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageModerationResult 图片审核结果。
type ImageModerationResult struct {

	// 审核情况。
	Suggestion *string `json:"suggestion,omitempty"`

	CategorySuggestions *CategorySuggestions `json:"category_suggestions,omitempty"`
}

func (o ImageModerationResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageModerationResult struct{}"
	}

	return strings.Join([]string{"ImageModerationResult", string(data)}, " ")
}
