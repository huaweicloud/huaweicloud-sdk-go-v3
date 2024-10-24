package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AgentImportParamNew struct {

	// 机器登录密码。
	Password string `json:"password"`

	// agent唯一值，重复导入时需要传递。
	AgentId *string `json:"agent_id,omitempty"`

	// 机器IP。
	InnerIp string `json:"inner_ip"`

	// 机器登录端口，默认22。
	Port int32 `json:"port"`

	// 机器ssh账号。
	Account string `json:"account"`

	// 机器操作系统类型。
	OsType string `json:"os_type"`

	// 机器所属VPC ID。
	VpcId *string `json:"vpc_id,omitempty"`

	// 外来唯一标识，COC用。
	CocCmdbId *string `json:"coc_cmdb_id,omitempty"`
}

func (o AgentImportParamNew) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentImportParamNew struct{}"
	}

	return strings.Join([]string{"AgentImportParamNew", string(data)}, " ")
}
