package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutopilotReleaseResponse Response Object
type ShowAutopilotReleaseResponse struct {

	// 模板名称
	ChartName *string `json:"chart_name,omitempty"`

	// 是否公开模板
	ChartPublic *bool `json:"chart_public,omitempty"`

	// 模板版本
	ChartVersion *string `json:"chart_version,omitempty"`

	// 集群ID
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 创建时间
	CreateAt *string `json:"create_at,omitempty"`

	// 模板实例描述
	Description *string `json:"description,omitempty"`

	// 模板实例名称
	Name *string `json:"name,omitempty"`

	// 模板实例所在的命名空间
	Namespace *string `json:"namespace,omitempty"`

	// 模板实例参数
	Parameters *string `json:"parameters,omitempty"`

	// 模板实例需要的资源
	Resources *string `json:"resources,omitempty"`

	// 模板实例状态 - DEPLOYED：已部署，表示模板实例处于正常状态。 - DELETED：已删除，表示模板实例已经被删除。 - FAILED：失败，表示模板实例部署失败。 - DELETING：删除中，表示模板实例正处于删除过程中。 - PENDING_INSTALL：待安装，表示模板正在等待安装。 - PENDING_UPGRADE：待升级，表示模板正在等待升级。 - PENDING_ROLLBACK：待回滚，表示模板正在等待回滚。 - UNKNOWN：未知，表示模板状态异常，可尝试手动删除后重新安装。
	Status *string `json:"status,omitempty"`

	// 模板实例状态描述
	StatusDescription *string `json:"status_description,omitempty"`

	// 更新时间
	UpdateAt *string `json:"update_at,omitempty"`

	// 模板实例的值
	Values *string `json:"values,omitempty"`

	// 模板实例版本
	Version        *int32 `json:"version,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowAutopilotReleaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutopilotReleaseResponse struct{}"
	}

	return strings.Join([]string{"ShowAutopilotReleaseResponse", string(data)}, " ")
}
