package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 标签的多种语言输出。
type ImageMediaTaggingDetItemBodyI18nTag struct {

	// 中文标签。
	Zh *string `json:"zh,omitempty"`

	// 英文标签。
	En *string `json:"en,omitempty"`
}

func (o ImageMediaTaggingDetItemBodyI18nTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageMediaTaggingDetItemBodyI18nTag struct{}"
	}

	return strings.Join([]string{"ImageMediaTaggingDetItemBodyI18nTag", string(data)}, " ")
}
