package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 实例详情
type DatabaseDto struct {

	// 实例id
	Id *string `json:"id,omitempty"`

	// 实例名称
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 模板id
	TemplateId *string `json:"template_id,omitempty"`

	// 模板名称
	TemplateName *string `json:"template_name,omitempty"`

	// 创建者
	Creator *string `json:"creator,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 数据条目
	DataCount *int64 `json:"data_count,omitempty"`

	// 源项目名
	SourceProjectName *string `json:"source_project_name,omitempty"`

	// 源项目id
	SourceProjectId *string `json:"source_project_id,omitempty"`

	// 源实例id
	SourceId *string `json:"source_id,omitempty"`

	// 是否为预置实例
	IsPrefab *bool `json:"is_prefab,omitempty"`
}

func (o DatabaseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseDto struct{}"
	}

	return strings.Join([]string{"DatabaseDto", string(data)}, " ")
}
