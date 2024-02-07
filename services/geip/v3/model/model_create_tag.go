package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTag 资源标签。
type CreateTag struct {

	// 标签键，最大长度128个unicode字符，格式为大小写字母，数字，中划线“-”，下划线“_”，中文。
	Key string `json:"key"`

	// 标签值，最大长度255个unicode字符，格式为大小写字母，数字，中划线“-”，下划线“_”，点“.”，中文。
	Value *string `json:"value,omitempty"`
}

func (o CreateTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTag struct{}"
	}

	return strings.Join([]string{"CreateTag", string(data)}, " ")
}
