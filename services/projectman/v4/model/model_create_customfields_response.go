package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateCustomfieldsResponse struct {
	// 字段选项

	Options *string `json:"options,omitempty"`
	// 系统字段

	Region *string `json:"region,omitempty"`
	// 字段ID

	Id *int32 `json:"id,omitempty"`
	// 字段ID

	Identifier *string `json:"identifier,omitempty"`
	// 项目ID

	ProjectId *int32 `json:"project_id,omitempty"`
	// 工作项类型id

	TrackerId *int32 `json:"tracker_id,omitempty"`
	// 系统字段名

	CustomField *string `json:"custom_field,omitempty"`
	// 字段类型

	Type *string `json:"type,omitempty"`
	// 字段名称

	Name *string `json:"name,omitempty"`
	// 系统字段

	Sort *int32 `json:"sort,omitempty"`
	// 字段描述

	Memo *string `json:"memo,omitempty"`
	// 创建时间

	Created *string `json:"created,omitempty"`
	// 修改时间

	Modified *string `json:"modified,omitempty"`
	// 是否被删除

	IsDelete       *bool `json:"is_delete,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CreateCustomfieldsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCustomfieldsResponse struct{}"
	}

	return strings.Join([]string{"CreateCustomfieldsResponse", string(data)}, " ")
}
