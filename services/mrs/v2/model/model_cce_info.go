package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CceInfo 容器信息。
type CceInfo struct {

	// pod名称。
	PodName *string `json:"pod_name,omitempty"`

	// pod的IP。
	PodIp *string `json:"pod_ip,omitempty"`

	// pod的ID。
	PodId *string `json:"pod_id,omitempty"`

	// 容器集群ID。对应CCE集群ID。
	CceClusterId *string `json:"cce_cluster_id,omitempty"`

	// 容器集群命名空间。
	CceNamespace *string `json:"cce_namespace,omitempty"`

	// cce节点IP。
	CceNodeIp *string `json:"cce_node_ip,omitempty"`
}

func (o CceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CceInfo struct{}"
	}

	return strings.Join([]string{"CceInfo", string(data)}, " ")
}
