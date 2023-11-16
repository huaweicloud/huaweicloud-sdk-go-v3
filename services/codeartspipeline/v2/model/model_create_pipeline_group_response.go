package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePipelineGroupResponse Response Object
type CreatePipelineGroupResponse struct {

	// 分组ID
	Id *string `json:"id,omitempty"`

	// 租户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 分组名
	Name *string `json:"name,omitempty"`

	// 父分组ID
	ParentId *string `json:"parent_id,omitempty"`

	// 分组路径ID
	PathId *string `json:"path_id,omitempty"`

	// 序号
	Ordinal *int32 `json:"ordinal,omitempty"`

	// 创建用户ID
	Creator *string `json:"creator,omitempty"`

	// 更新用户ID
	Updater *string `json:"updater,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 子分组列表
	Children       *[]PipelineGroupVo `json:"children,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o CreatePipelineGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePipelineGroupResponse struct{}"
	}

	return strings.Join([]string{"CreatePipelineGroupResponse", string(data)}, " ")
}
