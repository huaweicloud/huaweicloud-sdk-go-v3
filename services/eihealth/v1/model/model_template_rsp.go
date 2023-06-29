package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TemplateRsp 模板详情
type TemplateRsp struct {

	// 模板id
	Id *string `json:"id,omitempty"`

	// 模板名称
	Name *string `json:"name,omitempty"`

	// 模板描述
	Description *string `json:"description,omitempty"`

	// 来源项目名称
	SourceProjectName *string `json:"source_project_name,omitempty"`

	// 来源项目id
	SourceProjectId *string `json:"source_project_id,omitempty"`

	// 来源模板id
	SourceTemplateId *string `json:"source_template_id,omitempty"`

	// 创建者
	Creator *string `json:"creator,omitempty"`

	// 数据库列信息列表
	Columns *[]DatabaseColumnDto `json:"columns,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 主键
	PrimaryKey *string `json:"primary_key,omitempty"`

	// 是否是预置模板
	IsPrefab *bool `json:"is_prefab,omitempty"`
}

func (o TemplateRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateRsp struct{}"
	}

	return strings.Join([]string{"TemplateRsp", string(data)}, " ")
}
