package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 标签的多种语言输出。
type ImageTaggingItemBodyI18nTag struct {

	// 中文标签
	Zh *string `json:"zh,omitempty" xml:"zh"`

	// 英文标签
	En *string `json:"en,omitempty" xml:"en"`
}

func (o ImageTaggingItemBodyI18nTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageTaggingItemBodyI18nTag struct{}"
	}

	return strings.Join([]string{"ImageTaggingItemBodyI18nTag", string(data)}, " ")
}
