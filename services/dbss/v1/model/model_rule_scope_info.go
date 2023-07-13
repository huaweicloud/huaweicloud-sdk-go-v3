package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleScopeInfo struct {

	// 审计范围规则ID
	Id *string `json:"id,omitempty"`

	// 审计范围名称
	Name *string `json:"name,omitempty"`

	// 审计范围动作
	Action *string `json:"action,omitempty"`

	// 审计范围规则状态
	Status *string `json:"status,omitempty"`

	// 审计范围例外IP
	ExceptionIps *string `json:"exception_ips,omitempty"`

	// 审计范围规则源IP
	SourceIps *string `json:"source_ips,omitempty"`

	// 审计范围源端口
	SourcePorts *string `json:"source_ports,omitempty"`

	// 数据库ID
	DbIds *string `json:"db_ids,omitempty"`

	// 数据库名称
	DbNames *string `json:"db_names,omitempty"`

	// 数据库用户
	DbUsers *string `json:"db_users,omitempty"`

	// 是否全审计
	AllAudit *bool `json:"all_audit,omitempty"`
}

func (o RuleScopeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleScopeInfo struct{}"
	}

	return strings.Join([]string{"RuleScopeInfo", string(data)}, " ")
}
