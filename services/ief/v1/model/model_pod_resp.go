package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PodResp struct {

	// 应用实例uuid
	Id string `json:"id" xml:"id"`

	// 应用实例名称
	Name string `json:"name" xml:"name"`

	Configs *PodConfigs `json:"configs" xml:"configs"`

	Affinity *Affinity `json:"affinity" xml:"affinity"`

	// 应用实例init容器
	InitContainers *[]ContainerResp `json:"init_containers,omitempty" xml:"init_containers"`

	// 应用实例业务容器
	Containers []ContainerResp `json:"containers" xml:"containers"`

	// 应用实例所在节点
	NodeId string `json:"node_id" xml:"node_id"`

	// 应用ID
	DeploymentId string `json:"deployment_id" xml:"deployment_id"`

	// 项目ID
	ProjectId string `json:"project_id" xml:"project_id"`

	// 应用实例故障原因
	Reason string `json:"reason" xml:"reason"`

	// 应用实例故障详情
	Message string `json:"message" xml:"message"`

	// 应用实例创建时间
	CreatedAt string `json:"created_at" xml:"created_at"`

	// 应用实例更新时间
	UpdatedAt string `json:"updated_at" xml:"updated_at"`

	// 应用实例状态
	State string `json:"state" xml:"state"`
}

func (o PodResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PodResp struct{}"
	}

	return strings.Join([]string{"PodResp", string(data)}, " ")
}
