package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Project struct {

	// devcloud项目的数字id
	ProjectNumId *int32 `json:"project_num_id,omitempty"`

	// devcloud项目的32位id
	ProjectId *string `json:"project_id,omitempty"`

	// 项目名称
	Name *string `json:"name,omitempty"`

	// 项目描述
	Description *string `json:"description,omitempty"`

	// 项目创建时间
	CreatedTime *int64 `json:"created_time,omitempty"`

	// 项目更新时间
	UpdatedTime *int64 `json:"updated_time,omitempty"`

	// 项目代号
	ProjectCode *string `json:"project_code,omitempty"`

	// 区域region
	Region *string `json:"region,omitempty"`

	// 是否归档
	IsArchived *bool `json:"is_archived,omitempty"`

	// 项目类型
	Type *string `json:"type,omitempty"`

	Creator *User `json:"creator,omitempty"`
}

func (o Project) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Project struct{}"
	}

	return strings.Join([]string{"Project", string(data)}, " ")
}
