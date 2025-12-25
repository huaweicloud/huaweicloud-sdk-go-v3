package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModuleDetailInfo 模块详情
type ModuleDetailInfo struct {

	// 订阅包id
	CloudPackId *string `json:"cloud_pack_id,omitempty"`

	// 订阅包名称
	CloudPackName *string `json:"cloud_pack_name,omitempty"`

	// 订阅包版本
	CloudPackVersion *string `json:"cloud_pack_version,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 创建者ID
	CreatorId *string `json:"creator_id,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 英文描述
	EnDescription *string `json:"en_description,omitempty"`

	// 模块ID
	Id *string `json:"id,omitempty"`

	// 模块相关信息，当module为指标卡片时，items中的字段id也为指标id
	ModuleJson *string `json:"module_json,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 英文名称
	EnName *string `json:"en_name,omitempty"`

	// 租户ID
	ProjectId *string `json:"project_id,omitempty"`

	// 工作空间ID
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 模块缩略图
	Thumbnail *string `json:"thumbnail,omitempty"`

	// 模块类型,tab/section
	ModuleType *string `json:"module_type,omitempty"`

	// 模块标签
	Tag *string `json:"tag,omitempty"`

	// 是否为系统模块
	IsBuiltIn *bool `json:"is_built_in,omitempty"`

	// 数据查询方式
	DataQuery *string `json:"data_query,omitempty"`

	// BOA底座版本
	BoaVersion *string `json:"boa_version,omitempty"`

	// 安全云脑版本
	Version *string `json:"version,omitempty"`
}

func (o ModuleDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModuleDetailInfo struct{}"
	}

	return strings.Join([]string{"ModuleDetailInfo", string(data)}, " ")
}
