package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TemplateTagOptions 标签
type TemplateTagOptions struct {

	// 标签列表
	Tags *[]TemplateTag `json:"tags,omitempty"`
}

func (o TemplateTagOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateTagOptions struct{}"
	}

	return strings.Join([]string{"TemplateTagOptions", string(data)}, " ")
}
