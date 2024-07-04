package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComponentOverview struct {

	// 组件ID。
	Id *string `json:"id,omitempty"`

	// 应用ID。
	ApplicationId *string `json:"application_id,omitempty"`

	// 应用组件名称。
	Name *string `json:"name,omitempty"`

	// 项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	Runtime *RuntimeType `json:"runtime,omitempty"`

	Category *ComponentCategory `json:"category,omitempty"`

	SubCategory *ComponentSubCategory `json:"sub_category,omitempty"`

	// 组件描述。
	Description *string `json:"description,omitempty"`

	// 取值0或1。  0：表示正常状态。  1：表示正在删除。
	Status *int32 `json:"status,omitempty"`

	Source *SourceObject `json:"source,omitempty"`

	Build *BuildInfo `json:"build,omitempty"`

	// 创建人。
	Creator *string `json:"creator,omitempty"`

	// 创建时间。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 修改时间。
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 组件部署信息列表。
	Instances *[]ComponentInstanceOverView `json:"instances,omitempty"`
}

func (o ComponentOverview) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentOverview struct{}"
	}

	return strings.Join([]string{"ComponentOverview", string(data)}, " ")
}
