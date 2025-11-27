package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WebTamperEventContainerInfo **参数解释**: 网页防篡改事件信息对应的容器信息 **取值范围**: 不涉及
type WebTamperEventContainerInfo struct {

	// **参数解释**： 镜像名称 **取值范围**： 字符长度1-256位
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**： 镜像版本 **取值范围**： 字符长度1-256位
	ImageVersion *string `json:"image_version,omitempty"`

	// **参数解释**： 容器ID **取值范围**： 字符长度1-256位
	ContainerId *string `json:"container_id,omitempty"`

	// **参数解释**： 容器名称 **取值范围**： 字符长度1-256位
	ContainerName *string `json:"container_name,omitempty"`

	// **参数解释**： pod name **取值范围**： 字符长度1-256位
	PodName *string `json:"pod_name,omitempty"`

	// **参数解释**： pod ip **取值范围**： 字符长度1-256位
	PodIp *string `json:"pod_ip,omitempty"`

	// **参数解释**： 命名空间 **取值范围**： 字符长度1-256位
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**： 集群ID **取值范围**： 字符长度1-128位
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**： 集群名称 **取值范围**： 字符长度1-256位
	ClusterName *string `json:"cluster_name,omitempty"`
}

func (o WebTamperEventContainerInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebTamperEventContainerInfo struct{}"
	}

	return strings.Join([]string{"WebTamperEventContainerInfo", string(data)}, " ")
}
