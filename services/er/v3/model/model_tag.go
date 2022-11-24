package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 资源标签
type Tag struct {

	// 标签键，最大长度36个unicode字符，格式为大小写字母，数字，中划线“-”，下划线“_”，中文。
	Key *string `json:"key,omitempty"`

	// 标签值，最大长度43个unicode字符，格式为大小写字母，数字，中划线“-”，下划线“_”，点“.”，中文。
	Value *string `json:"value,omitempty"`
}

func (o Tag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tag struct{}"
	}

	return strings.Join([]string{"Tag", string(data)}, " ")
}
