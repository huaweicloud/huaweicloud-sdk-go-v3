package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDeploymentResponse struct {

	// 应用部署uuid
	Id *string `json:"id,omitempty" xml:"id"`

	// 应用部署名称，只允许英文小写字母、数字、中划线，最大长度32， 英文小写字母或数字开头和结尾
	Name *string `json:"name,omitempty" xml:"name"`

	// 应用部署总副本数
	Replicas *int32 `json:"replicas,omitempty" xml:"replicas"`

	// 应用部署正常副本数
	ReadyReplicas *int32 `json:"ready_replicas,omitempty" xml:"ready_replicas"`

	// 应用部署描述，最大长度255，不允许^ ~ # $ % & * < > ( ) [ ] { } ' \" \\
	Description *string `json:"description,omitempty" xml:"description"`

	// 应用部署到指定节点组，与node_ids二选一
	GroupId *string `json:"group_id,omitempty" xml:"group_id"`

	// 应用部署到指定节点，当前只支持一个边缘节点
	NodeIds *[]string `json:"node_ids,omitempty" xml:"node_ids"`

	// 节点属性
	Tags *[]Attributes `json:"tags,omitempty" xml:"tags"`

	// 应用部署版本
	ApiVersion *string `json:"api_version,omitempty" xml:"api_version"`

	// 应用部署来源：边缘市场（iem）或自定义（userdefined）
	Source *string `json:"source,omitempty" xml:"source"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty" xml:"project_id"`

	// 应用部署创建时间
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at"`

	// 应用部署更新时间
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at"`

	Template *PodRequest `json:"template,omitempty" xml:"template"`

	// 应用状态，仅包括冻结（FREEZE）、删除中（PENDING_DELETE）、删除失败（DELETE_FAILED），保留字段
	State *string `json:"state,omitempty" xml:"state"`

	// 预留字段
	SourceId *string `json:"source_id,omitempty" xml:"source_id"`

	Annotations    *Annotations `json:"annotations,omitempty" xml:"annotations"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowDeploymentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeploymentResponse struct{}"
	}

	return strings.Join([]string{"ShowDeploymentResponse", string(data)}, " ")
}
