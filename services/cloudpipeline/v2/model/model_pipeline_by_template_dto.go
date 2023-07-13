package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PipelineByTemplateDto struct {

	// 流水线名称
	Name *string `json:"name,omitempty"`

	// 流水线描述
	Description *string `json:"description,omitempty"`

	// 是否为变更流水线
	IsPublish *bool `json:"is_publish,omitempty"`

	// 流水线源
	Sources *[]CodeSource `json:"sources,omitempty"`
}

func (o PipelineByTemplateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineByTemplateDto struct{}"
	}

	return strings.Join([]string{"PipelineByTemplateDto", string(data)}, " ")
}
