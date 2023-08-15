package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageMediaTaggingItemBodyI18nType 标签类别的多种语言输出。
type ImageMediaTaggingItemBodyI18nType struct {

	// 中文标签类别。
	Zh *string `json:"zh,omitempty"`

	// 英文标签类别。
	En *string `json:"en,omitempty"`
}

func (o ImageMediaTaggingItemBodyI18nType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageMediaTaggingItemBodyI18nType struct{}"
	}

	return strings.Join([]string{"ImageMediaTaggingItemBodyI18nType", string(data)}, " ")
}
