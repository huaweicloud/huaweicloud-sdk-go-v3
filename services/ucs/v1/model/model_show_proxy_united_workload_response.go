package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProxyUnitedWorkloadResponse Response Object
type ShowProxyUnitedWorkloadResponse struct {

	// 资源类型，有以下取值： - Deployment：用于管理无状态应用 - Service：实现服务发现和负载均衡 - Ingress：管理对集群内服务的外部 HTTP/HTTPS 访问 - ConfigMap：用于存储非敏感的配置数据 - Secret：用于存储敏感信息 - Job：用于运行一次性任务的资源 - StatefulSet：用于管理有状态应用 - DaemonSet：用于在每个节点上运行一个 Pod 的资源 - PersistentVolumeClaim：用于请求和使用持久化存储资源的资源
	Kind *string `json:"kind,omitempty"`

	// API版本
	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *ObjectMeta `json:"metadata,omitempty"`

	// 要部署的应用模板
	Template *interface{} `json:"template,omitempty"`

	PropagationPolicy *PropagationPolicy `json:"propagationPolicy,omitempty"`

	OverridePolicy *OverridePolicy `json:"overridePolicy,omitempty"`

	ResourceBinding *ResourceBinding `json:"resourceBinding,omitempty"`
	HttpStatusCode  int              `json:"-"`
}

func (o ShowProxyUnitedWorkloadResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProxyUnitedWorkloadResponse struct{}"
	}

	return strings.Join([]string{"ShowProxyUnitedWorkloadResponse", string(data)}, " ")
}
