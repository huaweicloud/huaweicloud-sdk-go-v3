package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListQuotaResult struct {

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 企业项目名称。
	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`

	// EPS实例资源配额数量，值为-1时表示配额无限制。
	InstanceEpsQuota *int32 `json:"instance_eps_quota,omitempty"`

	// EPS计算资源配额数量，值为-1时表示配额无限制。
	VcpusEpsQuota *int32 `json:"vcpus_eps_quota,omitempty"`

	// EPS内存资源配额量，单位为GB，值为-1时表示配额无限制。
	RamEpsQuota *int32 `json:"ram_eps_quota,omitempty"`

	// EPS磁盘资源配额量，单位为GB，值为-1时表示配额无限制。
	VolumeEpsQuota *int32 `json:"volume_eps_quota,omitempty"`

	// EPS实例使用数量。
	InstanceUsed *int32 `json:"instance_used,omitempty"`

	// EPS计算资源使用数量。
	VcpusUsed *int32 `json:"vcpus_used,omitempty"`

	// EPS内存使用配额量，单位为GB。
	RamUsed *int32 `json:"ram_used,omitempty"`

	// EPS磁盘使用配额量，单位为GB。
	VolumeUsed *int32 `json:"volume_used,omitempty"`
}

func (o ListQuotaResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuotaResult struct{}"
	}

	return strings.Join([]string{"ListQuotaResult", string(data)}, " ")
}
