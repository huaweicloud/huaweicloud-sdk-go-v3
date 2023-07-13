package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceResponse Response Object
type ShowInstanceResponse struct {

	// 实例id
	Id *string `json:"id,omitempty"`

	// 实例名称
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	Template *TemplateRsp `json:"template,omitempty"`

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
	IsPrefab       *bool `json:"is_prefab,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceResponse", string(data)}, " ")
}
