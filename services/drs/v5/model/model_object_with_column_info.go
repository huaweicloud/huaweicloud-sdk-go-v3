package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ObjectWithColumnInfo 和列信息相关的对象
type ObjectWithColumnInfo struct {

	// 本节点id
	Id *string `json:"id,omitempty"`

	// 父节点id
	ParentId *string `json:"parent_id,omitempty"`

	// 节点类型
	Type *string `json:"type,omitempty"`

	// 节点名称
	Name *string `json:"name,omitempty"`

	// 节点别名
	AliasName *string `json:"alias_name,omitempty"`

	// 提示信息，例如提示库下表过多
	Notices *[]string `json:"notices,omitempty"`

	// 扩展信息，例如提示有无数据,结构是否在目标库中存在
	ExtendInfo *string `json:"extend_info,omitempty"`

	// 是否支持展开查询
	IsSupportExpand *bool `json:"is_support_expand,omitempty"`

	// 是否有列信息
	HasColumnInfo *bool `json:"has_column_info,omitempty"`

	// 是否预置
	IsPreset *bool `json:"is_preset,omitempty"`

	// token个数
	TokenCount *string `json:"token_count,omitempty"`

	// 是否已经下发给node
	IsSent *bool `json:"is_sent,omitempty"`

	// 下发给node的别名
	SentAliasName *string `json:"sent_alias_name,omitempty"`
}

func (o ObjectWithColumnInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObjectWithColumnInfo struct{}"
	}

	return strings.Join([]string{"ObjectWithColumnInfo", string(data)}, " ")
}
