package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateCustomfieldsResponse struct {

	// 字段选项
	Options *string `json:"options,omitempty" xml:"options"`

	// 系统字段
	Region *string `json:"region,omitempty" xml:"region"`

	// 字段ID
	Id *int32 `json:"id,omitempty" xml:"id"`

	// 字段ID
	Identifier *string `json:"identifier,omitempty" xml:"identifier"`

	// 项目ID
	ProjectId *int32 `json:"project_id,omitempty" xml:"project_id"`

	// 工作项类型id 2任务/Task,3缺陷/Bug,5Epic,6Feature,7Story
	TrackerId *int32 `json:"tracker_id,omitempty" xml:"tracker_id"`

	// 系统字段名
	CustomField *string `json:"custom_field,omitempty" xml:"custom_field"`

	// 字段类型
	Type *string `json:"type,omitempty" xml:"type"`

	// 字段名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 系统字段
	Sort *int32 `json:"sort,omitempty" xml:"sort"`

	// 字段描述
	Memo *string `json:"memo,omitempty" xml:"memo"`

	// 创建时间
	Created *string `json:"created,omitempty" xml:"created"`

	// 修改时间
	Modified *string `json:"modified,omitempty" xml:"modified"`

	// 是否被删除
	IsDelete       *bool `json:"is_delete,omitempty" xml:"is_delete"`
	HttpStatusCode int   `json:"-"`
}

func (o CreateCustomfieldsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCustomfieldsResponse struct{}"
	}

	return strings.Join([]string{"CreateCustomfieldsResponse", string(data)}, " ")
}
