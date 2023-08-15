package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageTaggingItemBodyI18nTag 标签的多种语言输出。
type ImageTaggingItemBodyI18nTag struct {

	// 中文标签
	Zh *string `json:"zh,omitempty"`

	// 英文标签
	En *string `json:"en,omitempty"`
}

func (o ImageTaggingItemBodyI18nTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageTaggingItemBodyI18nTag struct{}"
	}

	return strings.Join([]string{"ImageTaggingItemBodyI18nTag", string(data)}, " ")
}
