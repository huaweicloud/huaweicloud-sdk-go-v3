package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAppResponse Response Object
type ShowAppResponse struct {

	// 应用id
	Id *string `json:"id,omitempty"`

	// 应用名称
	Name *string `json:"name,omitempty"`

	// 应用版本
	Version *string `json:"version,omitempty"`

	// 应用短描述
	Summary *string `json:"summary,omitempty"`

	// 应用描述
	Description *string `json:"description,omitempty"`

	// 应用标签
	Labels *[]string `json:"labels,omitempty"`

	// 应用镜像
	Image *string `json:"image,omitempty"`

	// 应用命令
	Commands *[]string `json:"commands,omitempty"`

	Resources *ResourceDto `json:"resources,omitempty"`

	// 应用的输入参数信息
	Inputs *[]AppInputParameterDto `json:"inputs,omitempty"`

	// 应用的输出参数信息
	Outputs *[]AppOutputParameterDto `json:"outputs,omitempty"`

	// 创建应用时间
	CreateTime *string `json:"create_time,omitempty"`

	// 更新应用时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 创建应用的用户名
	UserName *string `json:"user_name,omitempty"`

	// 源项目名称
	SourceProjectName *string `json:"source_project_name,omitempty"`

	// 源资源id
	SourceResourceId *string `json:"source_resource_id,omitempty"`

	// 节点标签
	NodeLabels *[]string `json:"node_labels,omitempty"`

	// 图标base64编码
	Icon           *string `json:"icon,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppResponse struct{}"
	}

	return strings.Join([]string{"ShowAppResponse", string(data)}, " ")
}
