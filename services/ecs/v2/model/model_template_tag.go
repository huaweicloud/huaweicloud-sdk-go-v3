package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TemplateTag 标签列表
type TemplateTag struct {

	// 标签键
	Key *string `json:"key,omitempty"`

	// 标签值
	Value *string `json:"value,omitempty"`
}

func (o TemplateTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateTag struct{}"
	}

	return strings.Join([]string{"TemplateTag", string(data)}, " ")
}
