package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AgentInfoResult struct {

	// 机器 IP。
	InnerIp *string `json:"inner_ip,omitempty"`

	// 机器导入IP。
	ImportIp *string `json:"import_ip,omitempty"`

	// agent id。
	AgentId *string `json:"agent_id,omitempty"`

	// 主机名。
	HostName *string `json:"host_name,omitempty"`

	// 操作系统。
	OsType *string `json:"os_type,omitempty"`

	// UniAgent状态。
	AgentState *string `json:"agent_state,omitempty"`

	// 所属project ID。
	ProjectId *string `json:"project_id,omitempty"`

	// UniAgent版本。
	Version *string `json:"version,omitempty"`

	// 是否华为云机器。
	IsHwCloudHost *string `json:"is_hw_cloud_host,omitempty"`

	// VPC ID。
	VpcId *string `json:"vpc_id,omitempty"`

	// CMDB ID.
	CmdbId *string `json:"cmdb_id,omitempty"`

	// ECS ID。唯一值.
	EcsId *string `json:"ecs_id,omitempty"`

	// 机器所属domain ID。
	DomainId *string `json:"domain_id,omitempty"`
}

func (o AgentInfoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentInfoResult struct{}"
	}

	return strings.Join([]string{"AgentInfoResult", string(data)}, " ")
}
