package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModuleCreateRequestPojo struct {

	// 模块ID
	Id *string `json:"id,omitempty"`

	// 名称
	Name string `json:"name"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 租户ID
	ProjectId *string `json:"project_id,omitempty"`

	// 工作空间ID
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 模块相关信息
	ModuleJson *string `json:"module_json,omitempty"`

	// 模块类型,tab/section
	ModuleType *string `json:"module_type,omitempty"`

	// section类模块关联的指标id
	MetricIds *string `json:"metric_ids,omitempty"`

	// 模块缩略图
	Thumbnail *string `json:"thumbnail,omitempty"`

	// 数据查询方式
	DataQuery *string `json:"data_query,omitempty"`

	// BOA底座版本
	BoaVersion *string `json:"boa_version,omitempty"`
}

func (o ModuleCreateRequestPojo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModuleCreateRequestPojo struct{}"
	}

	return strings.Join([]string{"ModuleCreateRequestPojo", string(data)}, " ")
}
