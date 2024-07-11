package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AppGroupsEntity 分组信息
type AppGroupsEntity struct {

	// 分组id
	Id *string `json:"id,omitempty"`

	// 分组名称
	Name *string `json:"name,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 分组路径
	Path *string `json:"path,omitempty"`

	// 父分组id，首层为null
	ParentId *string `json:"parent_id,omitempty"`

	// 分组排序字段
	Ordinal *int32 `json:"ordinal,omitempty"`

	// 分组创建者用户id
	CreateUserId *string `json:"create_user_id,omitempty"`

	// 最近一次更新该分组用户id
	LastUpdateUserId *string `json:"last_update_user_id,omitempty"`

	// 该分组应用总数
	Count *int32 `json:"count,omitempty"`

	// 子分组列表
	Children *[]AppGroupsEntity `json:"children,omitempty"`
}

func (o AppGroupsEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppGroupsEntity struct{}"
	}

	return strings.Join([]string{"AppGroupsEntity", string(data)}, " ")
}
