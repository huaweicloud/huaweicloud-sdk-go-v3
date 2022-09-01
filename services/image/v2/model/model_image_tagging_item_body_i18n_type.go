package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 标签类别的多种语言输出。
type ImageTaggingItemBodyI18nType struct {

	// 中文标签类别
	Zh *string `json:"zh,omitempty" xml:"zh"`

	// 英文标签类别
	En *string `json:"en,omitempty" xml:"en"`
}

func (o ImageTaggingItemBodyI18nType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageTaggingItemBodyI18nType struct{}"
	}

	return strings.Join([]string{"ImageTaggingItemBodyI18nType", string(data)}, " ")
}
