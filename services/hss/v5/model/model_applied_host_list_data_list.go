package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AppliedHostListDataList struct {

	// **参数解释**： 服务器（主机）的唯一标识ID **取值范围**： 字符长度1-64位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**： 服务器私有IP **取值范围**： 字符长度1-128位
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**: 主机上安装的杀毒Agent的唯一标识ID，用于关联主机与杀毒服务 **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	AgentId *string `json:"agent_id,omitempty"`

	// 冲突的端口
	ConflictPort *[]int32 `json:"conflict_port,omitempty"`

	// 应用端口
	AppliedPort *[]int32 `json:"applied_port,omitempty"`
}

func (o AppliedHostListDataList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppliedHostListDataList struct{}"
	}

	return strings.Join([]string{"AppliedHostListDataList", string(data)}, " ")
}
