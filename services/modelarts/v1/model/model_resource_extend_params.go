package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceExtendParams **参数解释**：自定义配置。 **约束限制**：不涉及。
type ResourceExtendParams struct {

	// **参数解释**：节点的容器镜像空间大小。 **取值范围**：可选值如下： - 指定大小：dockerBaseSize的大小范围默认为50-500，但实际上限受到节点的容器数据盘大小约束。 - 不限制：dockerBaseSize=\\\"0\\\"，表示不限制容器镜像空间大小。
	DockerBaseSize *string `json:"dockerBaseSize,omitempty"`

	// **参数解释**：安装后执行脚本，输入的值需要经过Base64编码。 **取值范围**：不涉及。
	PostInstall *string `json:"postInstall,omitempty"`

	// **参数解释**：容器运行时。 **取值范围**：可选值如下： - docker：容器运行时，是目前最常用的容器化引擎，基于容器镜像创建和管理容器实例。 - containerd：工业级的容器运行时，专注于容器的生命周期管理，是 Docker 底层核心组件之一，也可独立部署使用。
	Runtime *string `json:"runtime,omitempty"`

	// **参数解释**：存量节点k8s标签更新策略，值为空时默认更新存量节点。 **取值范围**：可选值如下： - refresh：更新。 - ignore：不更新。
	LabelPolicyOnExistingNodes *string `json:"labelPolicyOnExistingNodes,omitempty"`

	// **参数解释**：存量节点k8s污点更新策略，值为空时默认更新存量节点。 **取值范围**：可选值如下： - refresh：更新。 - ignore：不更新。
	TaintPolicyOnExistingNodes *string `json:"taintPolicyOnExistingNodes,omitempty"`

	// **参数解释**：存量节点资源标签更新策略，值为空时默认更新存量节点。 **取值范围**：可选值如下： - refresh：更新。 - ignore：不更新。
	TagPolicyOnExistingNodes *string `json:"tagPolicyOnExistingNodes,omitempty"`

	// **参数解释**：跨物理集群之间进行参数面数据传输使用的子网id。不可与节点子网和容器子网重复。 **取值范围**：不涉及。
	XParameterPlaneSubnet *string `json:"XParameterPlaneSubnet,omitempty"`

	// **参数解释**：用户指定的节点池名称。最小长度为2，最大长度为50的小写字母、中划线-、数字组成，由小写字母开头，不能以-，-default结尾。 **取值范围**：不涉及
	NodePoolName *string `json:"nodePoolName,omitempty"`
}

func (o ResourceExtendParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceExtendParams struct{}"
	}

	return strings.Join([]string{"ResourceExtendParams", string(data)}, " ")
}
