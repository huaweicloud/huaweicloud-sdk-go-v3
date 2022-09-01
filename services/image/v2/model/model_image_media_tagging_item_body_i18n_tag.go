package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 标签的多种语言输出。
type ImageMediaTaggingItemBodyI18nTag struct {

	// 中文标签。
	Zh *string `json:"zh,omitempty" xml:"zh"`

	// 英文标签。
	En *string `json:"en,omitempty" xml:"en"`
}

func (o ImageMediaTaggingItemBodyI18nTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageMediaTaggingItemBodyI18nTag struct{}"
	}

	return strings.Join([]string{"ImageMediaTaggingItemBodyI18nTag", string(data)}, " ")
}