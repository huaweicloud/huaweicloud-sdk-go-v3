package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobGroupTreeResponseBody 分组树
type JobGroupTreeResponseBody struct {

	// 分组编号
	Id *string `json:"id,omitempty"`

	// 租户ID
	DomainId *string `json:"domain_id,omitempty"`

	// CodeArts项目ID。构建任务所在项目的ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 分组名称
	Name *string `json:"name,omitempty"`

	// 父分组编号
	ParentId *string `json:"parent_id,omitempty"`

	// 分组路径
	PathId *string `json:"path_id,omitempty"`

	// 序数
	Ordinal *int32 `json:"ordinal,omitempty"`

	// 创建用户
	CreateUser *string `json:"create_user,omitempty"`

	// 更新用户
	UpdateUser *string `json:"update_user,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 子分组
	Children *[]JobGroupTreeResponseBody `json:"children,omitempty"`
}

func (o JobGroupTreeResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobGroupTreeResponseBody struct{}"
	}

	return strings.Join([]string{"JobGroupTreeResponseBody", string(data)}, " ")
}
