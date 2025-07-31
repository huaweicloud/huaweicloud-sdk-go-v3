package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DaemonSetInfo daemonset基本信息
type DaemonSetInfo struct {

	// daemonset名称
	Name *string `json:"name,omitempty"`

	// 命名空间名称
	NamespaceName *string `json:"namespace_name,omitempty"`

	// 集群id
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群类型，包含以下几种： -k8s 原生集群 -cce CCE集群 -ali 阿里云集群 -tencent 腾讯云集群 -azure 微软云集群 -aws 亚马逊集群 -self_built_hw 华为云自建集群 -self_built_idc IDC自建集群
	ClusterType *string `json:"cluster_type,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 状态，包含以下几种 -Running：正常运行 -Failed：存在异常
	Status *string `json:"status,omitempty"`

	// 实例个数
	PodsNum *int32 `json:"pods_num,omitempty"`

	// 镜像名称
	ImageName *string `json:"image_name,omitempty"`

	// 标签
	MatchLabels *[]LabelInfo `json:"match_labels,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`
}

func (o DaemonSetInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DaemonSetInfo struct{}"
	}

	return strings.Join([]string{"DaemonSetInfo", string(data)}, " ")
}
