package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TemplateListV4ResponseBodyTemplates struct {

	// 模板id
	Id *int32 `json:"id,omitempty"`

	// 项目id
	ProjectId *int32 `json:"project_id,omitempty"`

	// 工作项类型id
	TrackerId *int32 `json:"tracker_id,omitempty"`

	// 工作项详情模板描述内容
	Description *string `json:"description,omitempty"`

	// 模板配置
	IssueFieldConfig *string `json:"issue_field_config,omitempty"`
}

func (o TemplateListV4ResponseBodyTemplates) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateListV4ResponseBodyTemplates struct{}"
	}

	return strings.Join([]string{"TemplateListV4ResponseBodyTemplates", string(data)}, " ")
}
