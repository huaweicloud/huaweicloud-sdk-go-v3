package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterContainerLogResponseInfo k8s集群容器日志信息
type ClusterContainerLogResponseInfo struct {

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 集群id
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群类型，包含以下几种： - cce: cce集群 - ali: 阿里云集群 - tencent: 腾讯云集群 - azure: 微软云集群 - aws: 亚马逊集群 - self_built_hw: 华为云自建集群 - self_built_idc: IDC自建集群
	ClusterType *string `json:"cluster_type,omitempty"`

	// 日志产生的时间
	Time *int64 `json:"time,omitempty"`

	// 容器日志所属的命名空间
	Namespace *string `json:"namespace,omitempty"`

	// 产生日志的容器所属pod的名称
	PodName *string `json:"pod_name,omitempty"`

	// 产生日志的容器所属pod的id
	PodId *string `json:"pod_id,omitempty"`

	// 产生日志的容器所属pod的ip
	PodIp *string `json:"pod_ip,omitempty"`

	// 产生日志的容器所在的主机ip
	HostIp *string `json:"host_ip,omitempty"`

	// 产生日志的容器名称
	ContainerName *string `json:"container_name,omitempty"`

	// 产生日志的容器id
	ContainerId *string `json:"container_id,omitempty"`

	// 日志的内容
	Content *string `json:"content,omitempty"`

	// cce集群容器日志的行号
	LineNum *string `json:"line_num,omitempty"`
}

func (o ClusterContainerLogResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterContainerLogResponseInfo struct{}"
	}

	return strings.Join([]string{"ClusterContainerLogResponseInfo", string(data)}, " ")
}
