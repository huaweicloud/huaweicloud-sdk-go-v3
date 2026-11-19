package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterInfo 集群安装agent信息
type ClusterInfo struct {

	// **参数解释**: 所属企业项目名称 **取值范围**: 字符长度0-65535位
	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`

	// **参数解释**: 集群版本 **取值范围**: 字符长度0-256位
	ClusterVersion *string `json:"cluster_version,omitempty"`

	// **参数解释**: 集群id **取值范围**: 字符长度0-64位
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**: 所属集群 **取值范围**: 字符长度0-256位
	ClusterName *string `json:"cluster_name,omitempty"`

	// **参数解释**： 集群状态，从CCE同步状态 **取值范围**： - workload：工作负载。 - cluster：集群。
	ClusterStatus *string `json:"cluster_status,omitempty"`

	// **参数解释**: 已安装实例个数 **取值范围**: 最小值0，最大值2147483647
	InstalledPodNum *int32 `json:"installed_pod_num,omitempty"`

	// **参数解释**: 全部实例个数 **取值范围**: 最小值0，最大值2147483647
	PodNum *int32 `json:"pod_num,omitempty"`

	// **参数解释**: 已防护实例数 **取值范围**: 最小值0，最大值2147483647
	ProtectPodNum *int32 `json:"protect_pod_num,omitempty"`

	// **参数解释**: 创建时间 **取值范围**: 最小值0，最大值9223372036854775807
	CreateTime *int64 `json:"create_time,omitempty"`

	// **参数解释**: 防护状态 **取值范围**: 字符长度0-65535位
	ProtectStatus *string `json:"protect_status,omitempty"`
}

func (o ClusterInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterInfo struct{}"
	}

	return strings.Join([]string{"ClusterInfo", string(data)}, " ")
}
