package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceStatisticsResponse Response Object
type ListInstanceStatisticsResponse struct {

	// 已用存储空间
	StorageUsed *int64 `json:"storage_used,omitempty"`

	// 命名空间的总数
	TotalNamespaceCount *int64 `json:"total_namespace_count,omitempty"`

	// 镜像的总数
	TotalImageCount *int64 `json:"total_image_count,omitempty"`

	// 命名空间的配额
	NamespaceQuota *int64 `json:"namespace_quota,omitempty"`

	// 镜像的配额
	ImageRepoQuota *int64 `json:"image_repo_quota,omitempty"`

	// 镜像同步策略的配额
	ReplicaPolicyQuota *int64 `json:"replica_policy_quota,omitempty"`

	// 镜像老化策略的配额
	RetentionPolicyQuota *int64 `json:"retention_policy_quota,omitempty"`

	// 触发器的配额
	NotifyPolicyQuota *int64 `json:"notify_policy_quota,omitempty"`

	// 镜像同步的目标仓库配额
	ReplicaRegistryQuota *int64 `json:"replica_registry_quota,omitempty"`

	// 镜像签名策略的配额
	SignPolicyQuota *int64 `json:"sign_policy_quota,omitempty"`

	// 镜像同步策略总数
	ReplicaPolicyCount *int64 `json:"replica_policy_count,omitempty"`

	// 镜像老化策略的总数
	RetentionPolicyCount *int64 `json:"retention_policy_count,omitempty"`

	// 触发器的总数
	NotifyPolicyCount *int64 `json:"notify_policy_count,omitempty"`

	// 镜像同步策略的目标仓库总数
	ReplicaRegistryCount *int64 `json:"replica_registry_count,omitempty"`

	// 内网访问控制的配额
	IntranetEndpointQuota *int32 `json:"intranet_endpoint_quota,omitempty"`

	// 内网访问控制的总数
	IntranetEndpointCount *int32 `json:"intranet_endpoint_count,omitempty"`

	// 长期登录指令的配额
	LongTermQuota *int32 `json:"long_term_quota,omitempty"`

	// 镜像签名策略的总数
	SignPolicyCount *int64 `json:"sign_policy_count,omitempty"`
	HttpStatusCode  int    `json:"-"`
}

func (o ListInstanceStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceStatisticsResponse", string(data)}, " ")
}
