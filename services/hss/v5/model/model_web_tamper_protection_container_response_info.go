package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WebTamperProtectionContainerResponseInfo **参数解释**: 防护配置关联的容器信息 **取值范围**: 不涉及
type WebTamperProtectionContainerResponseInfo struct {

	// description: |   **参数解释**:   防护状态   **取值范围**:   - protected：防护中   - partial_protected：部分防护   - protect_failed：防护失败
	Status *string `json:"status,omitempty"`

	// **参数解释**： 容器ID **取值范围**： 字符长度1-256位
	ContainerId *string `json:"container_id,omitempty"`

	// **参数解释**： 容器名称 **取值范围**： 字符长度1-256位
	ContainerName *string `json:"container_name,omitempty"`

	// **参数解释**： 镜像名称 **取值范围**： 字符长度1-256位
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**： 镜像版本 **取值范围**： 字符长度1-256位
	ImageVersion *string `json:"image_version,omitempty"`

	// **参数解释**： pod name **取值范围**： 字符长度1-256位
	PodName *string `json:"pod_name,omitempty"`

	// **参数解释**： pod ip **取值范围**： 字符长度1-256位
	PodIp *string `json:"pod_ip,omitempty"`

	// **参数解释**： 命名空间 **取值范围**： 字符长度1-256位
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**: 防护总目录数量 **取值范围**: 最小值0，最大值2147483647
	AllDirectoryNums *int32 `json:"all_directory_nums,omitempty"`

	// **参数解释**: 防护成功数量 **取值范围**: 最小值0，最大值2147483647
	ProtectedDirectoryNums *int32 `json:"protected_directory_nums,omitempty"`

	HostInfo *WebTamperEventHostInfo `json:"host_info,omitempty"`

	ClusterInfo *WebTamperEventClusterInfo `json:"cluster_info,omitempty"`
}

func (o WebTamperProtectionContainerResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebTamperProtectionContainerResponseInfo struct{}"
	}

	return strings.Join([]string{"WebTamperProtectionContainerResponseInfo", string(data)}, " ")
}
