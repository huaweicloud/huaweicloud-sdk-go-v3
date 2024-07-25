package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkspacesRequest Request Object
type ListWorkspacesRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// 偏移量 指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，
	Offset float32 `json:"offset"`

	// 每页显示个数
	Limit float32 `json:"limit"`

	// 区域id
	RegionId *string `json:"region_id,omitempty"`

	// 名称查询
	Name *string `json:"name,omitempty"`

	// 描述查询
	Description *string `json:"description,omitempty"`

	// 视图绑定的空间id
	ViewBindId *string `json:"view_bind_id,omitempty"`

	// 视图绑定的空间名称
	ViewBindName *string `json:"view_bind_name,omitempty"`

	// 创建时间开始，例如2024-04-26T16:08:09Z+0800
	CreateTimeStart *string `json:"create_time_start,omitempty"`

	// 创建时间结束，例如2024-04-2T16:08:09Z+0800
	CreateTimeEnd *string `json:"create_time_end,omitempty"`

	// 是否查询视图 true or false
	IsView *bool `json:"is_view,omitempty"`

	// 工作空间id数组，英文逗号分隔
	Ids *string `json:"ids,omitempty"`

	// 普通项目的项目id
	NormalProjectId *string `json:"normal_project_id,omitempty"`

	// 企业项目的项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListWorkspacesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkspacesRequest struct{}"
	}

	return strings.Join([]string{"ListWorkspacesRequest", string(data)}, " ")
}
