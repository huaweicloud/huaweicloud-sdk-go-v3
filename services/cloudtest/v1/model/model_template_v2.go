package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TemplateV2 struct {

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 创建人名称
	CreatorName *string `json:"creator_name,omitempty"`

	// 创建人工号
	CreatorNum *string `json:"creator_num,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// id 主键
	Id *string `json:"id,omitempty"`

	// 是否默认
	IsDefault *string `json:"is_default,omitempty"`

	// 脑图json
	Mindmap *string `json:"mindmap,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 编辑权限
	EditPermission *bool `json:"editPermission,omitempty"`

	// 删除权限
	DeletePermission *bool `json:"deletePermission,omitempty"`
}

func (o TemplateV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateV2 struct{}"
	}

	return strings.Join([]string{"TemplateV2", string(data)}, " ")
}
