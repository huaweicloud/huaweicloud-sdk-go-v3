package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobGroupResponseBodyResult 结果
type JobGroupResponseBodyResult struct {

	// 主键id
	Id *string `json:"id,omitempty"`

	// 分组名称
	Name *string `json:"name,omitempty"`

	// CodeArts项目ID。构建任务所在项目的ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 父分组id
	ParentId *string `json:"parent_id,omitempty"`

	// 任务分组id
	GroupId *string `json:"group_id,omitempty"`

	// 租户id
	DomainId *string `json:"domain_id,omitempty"`

	// 分组的index
	Ordinal *string `json:"ordinal,omitempty"`

	// 创建者
	CreateUser *string `json:"create_user,omitempty"`

	// 修改者
	UpdateUser *string `json:"update_user,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 修改时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 分组地址id
	PathId *string `json:"path_id,omitempty"`
}

func (o JobGroupResponseBodyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobGroupResponseBodyResult struct{}"
	}

	return strings.Join([]string{"JobGroupResponseBodyResult", string(data)}, " ")
}
