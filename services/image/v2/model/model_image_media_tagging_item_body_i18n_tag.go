package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageMediaTaggingItemBodyI18nTag 标签的多种语言输出。
type ImageMediaTaggingItemBodyI18nTag struct {

	// 中文标签。
	Zh *string `json:"zh,omitempty"`

	// 英文标签。
	En *string `json:"en,omitempty"`
}

func (o ImageMediaTaggingItemBodyI18nTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageMediaTaggingItemBodyI18nTag struct{}"
	}

	return strings.Join([]string{"ImageMediaTaggingItemBodyI18nTag", string(data)}, " ")
}
