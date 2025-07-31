package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterItemResponseInfo 集群防护策略列表
type ClusterItemResponseInfo struct {

	// 集群id
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 集群命名空间
	ClusterNs *[]string `json:"cluster_ns,omitempty"`

	// 集群标签
	ClusterLabels *[]string `json:"cluster_labels,omitempty"`

	// 集群防护状态 | - \"0\" unprotected - \"1\" plugin error - \"2\" protected with policy - \"3\" deploy policy failed - \"4\" protected without policy - \"5\" uninstall failed - \"6\" uninstall - \"7\" installing
	ProtectStatus *string `json:"protect_status,omitempty"`
}

func (o ClusterItemResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterItemResponseInfo struct{}"
	}

	return strings.Join([]string{"ClusterItemResponseInfo", string(data)}, " ")
}
