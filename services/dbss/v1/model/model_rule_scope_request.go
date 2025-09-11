package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleScopeRequest struct {

	// 操作类型，多个用英文,分隔
	Action *string `json:"action,omitempty"`

	// 数据库ID，多个用英文逗号分隔，全部传[ALL]
	DbIds string `json:"db_ids"`

	// 数据库名称，多个用英文逗号分隔，全部传[全部数据库]
	DbNames string `json:"db_names"`

	// 数据库账号，多个用英文逗号分隔
	DbUsers *string `json:"db_users,omitempty"`

	// 例外IP，多个用英文逗号分隔
	ExceptionIps *string `json:"exception_ips,omitempty"`

	// 名称
	RuleName string `json:"rule_name"`

	// 源IP，多个用英文逗号分隔
	SourceIps *string `json:"source_ips,omitempty"`

	// 源端口，多个用英文逗号分隔
	SourcePorts *string `json:"source_ports,omitempty"`
}

func (o RuleScopeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleScopeRequest struct{}"
	}

	return strings.Join([]string{"RuleScopeRequest", string(data)}, " ")
}
