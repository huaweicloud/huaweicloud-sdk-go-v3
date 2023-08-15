package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Pod struct {
	Configs *PodConfig `json:"configs,omitempty"`

	// 部署失败的原因
	Reason *string `json:"reason,omitempty"`

	// 对应网卡地址
	HostIp *string `json:"host_ip,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 集群ID
	ClusterId *string `json:"cluster_id,omitempty"`

	// 更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 实例名字
	Name *string `json:"name,omitempty"`

	// 实例ID
	Id *string `json:"id,omitempty"`

	// 部署ID
	DeploymentId *string `json:"deployment_id,omitempty"`

	Affinity *PodAffinity `json:"affinity,omitempty"`

	// 应用部署信息
	Apps *[]AppDef `json:"apps,omitempty"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	// 状态，状态包括，Pending，表示挂起，Running表示pod已经被调到到某节点，Succeeded表示Pod已经被成功终止，Failed表示左右容器都已终止，Unkonwn表示无法取得Pod状态
	Status *string `json:"status,omitempty"`
}

func (o Pod) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Pod struct{}"
	}

	return strings.Join([]string{"Pod", string(data)}, " ")
}
