package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterResponseInfo 集群防护信息
type ClusterResponseInfo struct {

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 集群ID
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群版本
	ClusterVersion *string `json:"cluster_version,omitempty"`

	// 防护状态 - unprotected - plugin error - protected with policy - deploy policy failed - protected without policy - uninstall failed - uninstall
	ProtectStatus *string `json:"protect_status,omitempty"`

	// 策略数量
	PolicyNum *int32 `json:"policy_num,omitempty"`

	// 集群运行状态 - Available - Unavailable
	ClusterStatus *string `json:"cluster_status,omitempty"`

	// 集群类型，包含以下几种： - k8s：原生集群 - cce：CCE集群 - ali：阿里云集群 - tencent：腾讯云集群 - azure：微软云集群 - aws：亚马逊集群 - self_built_hw：华为云自建集群 - self_built_idc：IDC自建集群
	ClusterType *string `json:"cluster_type,omitempty"`
}

func (o ClusterResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterResponseInfo struct{}"
	}

	return strings.Join([]string{"ClusterResponseInfo", string(data)}, " ")
}
