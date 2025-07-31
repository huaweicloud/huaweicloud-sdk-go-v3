package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterAuditLogResponseInfo k8s集群审计日志信息
type ClusterAuditLogResponseInfo struct {

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 集群id
	ClusterId *string `json:"cluster_id,omitempty"`

	// 日志产生的时间
	Time *int64 `json:"time,omitempty"`

	// 审计日志的内容，json格式的字符串
	Content *string `json:"content,omitempty"`

	// 集群类型，包含以下几种： - cce: cce集群 - ali: 阿里云集群 - tencent: 腾讯云集群 - azure: 微软云集群 - aws: 亚马逊集群 - self_built_hw: 华为云自建集群 - self_built_idc: IDC自建集群
	ClusterType *string `json:"cluster_type,omitempty"`

	// 主机ID
	HostId *string `json:"host_id,omitempty"`

	// 主机名称
	HostName *string `json:"host_name,omitempty"`

	// 主机ip
	HostIp *string `json:"host_ip,omitempty"`

	// cce集群日志的行号
	LineNum *string `json:"line_num,omitempty"`
}

func (o ClusterAuditLogResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterAuditLogResponseInfo struct{}"
	}

	return strings.Join([]string{"ClusterAuditLogResponseInfo", string(data)}, " ")
}
