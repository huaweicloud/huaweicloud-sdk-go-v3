package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TemplateSimpleInfo struct {

	// 模板id。
	Id *string `json:"id,omitempty" xml:"id"`

	// 模板名。
	Title *string `json:"title,omitempty" xml:"title"`

	// 模板描述。
	Description *string `json:"description,omitempty" xml:"description"`
}

func (o TemplateSimpleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateSimpleInfo struct{}"
	}

	return strings.Join([]string{"TemplateSimpleInfo", string(data)}, " ")
}
