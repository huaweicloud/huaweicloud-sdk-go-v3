package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDeploymentPodsRequest struct {

	// 集群ID，查询部署在该节点组的应用列表，和node_id不可同时请求
	ClusterId *string `json:"cluster_id,omitempty"`

	// 节点ID，查询部署在该节点下的应用列表，和cluster_id不可同时请求
	NodeId *string `json:"node_id,omitempty"`

	// 平台提供者，分别为hilens及ief。当为hilens时，请求部署在hilens平台的相关数据
	Provider *string `json:"provider,omitempty"`

	// 应用部署ID
	DeploymentId *string `json:"deployment_id,omitempty"`

	// 工作空间ID，默认为注册账号/子账号的default工作空间。主账号默认ID为0，子账号默认空间id为该子账号user_id
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 查询的起始位置，取值范围为非负整数，默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量，取值范围1~100，默认为100
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowDeploymentPodsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeploymentPodsRequest struct{}"
	}

	return strings.Join([]string{"ShowDeploymentPodsRequest", string(data)}, " ")
}
