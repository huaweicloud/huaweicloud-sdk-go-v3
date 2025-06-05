package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AppDto 应用列表返回的单个应用对象。
type AppDto struct {

	// 应用ID。
	Id *string `json:"id,omitempty"`

	// 应用名称。
	Name *string `json:"name,omitempty"`

	// 应用所属空间ID。
	EihealthProjectId *string `json:"eihealth_project_id,omitempty"`

	// 应用所属空间名称。
	EihealthProjectName *string `json:"eihealth_project_name,omitempty"`

	// 应用版本。
	Version *string `json:"version,omitempty"`

	// 应用简述。
	Summary *string `json:"summary,omitempty"`

	// 应用描述。
	Description *string `json:"description,omitempty"`

	// 应用标签。
	Labels *[]string `json:"labels,omitempty"`

	// 创建应用时间。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新应用时间。
	UpdateTime *string `json:"update_time,omitempty"`

	// 创建应用的用户名。
	UserName *string `json:"user_name,omitempty"`

	// 源项目名称。
	SourceProjectName *string `json:"source_project_name,omitempty"`

	// 源资源ID。
	SourceResourceId *string `json:"source_resource_id,omitempty"`

	// 图标base64编码。
	Icon *string `json:"icon,omitempty"`
}

func (o AppDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppDto struct{}"
	}

	return strings.Join([]string{"AppDto", string(data)}, " ")
}
