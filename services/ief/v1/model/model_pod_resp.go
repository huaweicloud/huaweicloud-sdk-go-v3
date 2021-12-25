package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PodResp struct {
	// 应用实例uuid

	Id string `json:"id"`
	// 应用实例名称

	Name string `json:"name"`

	Configs *PodConfigs `json:"configs"`

	Affinity *Affinity `json:"affinity"`
	// 应用实例init容器

	InitContainers *[]ContainerResp `json:"init_containers,omitempty"`
	// 应用实例业务容器

	Containers []ContainerResp `json:"containers"`
	// 应用实例所在节点

	NodeId string `json:"node_id"`
	// 应用ID

	DeploymentId string `json:"deployment_id"`
	// 项目ID

	ProjectId string `json:"project_id"`
	// 应用实例故障原因

	Reason string `json:"reason"`
	// 应用实例故障详情

	Message string `json:"message"`
	// 应用实例创建时间

	CreatedAt string `json:"created_at"`
	// 应用实例更新时间

	UpdatedAt string `json:"updated_at"`
	// 应用实例状态

	State string `json:"state"`
}

func (o PodResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PodResp struct{}"
	}

	return strings.Join([]string{"PodResp", string(data)}, " ")
}
