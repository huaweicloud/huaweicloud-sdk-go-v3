package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateDeploymentResponse struct {
	// 应用部署uuid

	Id *string `json:"id,omitempty"`
	// 应用部署名称，只允许英文小写字母、数字、中划线，最大长度32, 英文小写字母或数字开头和结尾

	Name *string `json:"name,omitempty"`
	// 应用部署总副本数

	Replicas *int32 `json:"replicas,omitempty"`
	// 应用部署正常副本数

	ReadyReplicas *int32 `json:"ready_replicas,omitempty"`
	// 应用部署描述，最大长度255，不允许^ ~ # $ % & * < > ( ) [ ] { } ' \" \\

	Description *string `json:"description,omitempty"`
	// 应用部署到指定节点组，与node_ids二选一

	GroupId *string `json:"group_id,omitempty"`
	// 应用部署到指定节点，当前只支持一个边缘节点

	NodeIds *[]string `json:"node_ids,omitempty"`
	// 节点属性

	Tags *[]Attributes `json:"tags,omitempty"`
	// 应用部署版本

	ApiVersion *string `json:"api_version,omitempty"`
	// 应用部署来源: 边缘市场(iem) or 自定义(userdefined)

	Source *string `json:"source,omitempty"`
	// 项目ID

	ProjectId *string `json:"project_id,omitempty"`
	// 应用部署创建时间

	CreatedAt *string `json:"created_at,omitempty"`
	// 应用部署更新时间

	UpdatedAt *string `json:"updated_at,omitempty"`

	Template *PodRequest `json:"template,omitempty"`
	// 应用状态，仅包括冻结（FREEZE）、删除中（PENDING_DELETE）、删除失败（DELETE_FAILED），保留字段

	State *string `json:"state,omitempty"`
	// 预留字段

	SourceId       *string `json:"source_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDeploymentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDeploymentResponse struct{}"
	}

	return strings.Join([]string{"UpdateDeploymentResponse", string(data)}, " ")
}
