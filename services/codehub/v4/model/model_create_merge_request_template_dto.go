package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateMergeRequestTemplateDto struct {

	// **参数解释：** 模板名称
	TemplateName string `json:"template_name"`

	// **参数解释：** 合并请求标题（不自动提取合并请求标题时生效）
	MergeRequestTitle *string `json:"merge_request_title,omitempty"`

	// **参数解释：** 描述
	Description *string `json:"description,omitempty"`

	// **参数解释：** 自动提取合并请求标题 0：不提取 1：提取提交信息 2：提取e2e单标题
	AutoExtractMrTitle *int32 `json:"auto_extract_mr_title,omitempty"`

	// **参数解释：** 是否在标题中添加[WIP]
	IsWip *bool `json:"is_wip,omitempty"`

	// **参数解释：** 是否设置为默认模板
	IsDefault *bool `json:"is_default,omitempty"`
}

func (o CreateMergeRequestTemplateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMergeRequestTemplateDto struct{}"
	}

	return strings.Join([]string{"CreateMergeRequestTemplateDto", string(data)}, " ")
}
