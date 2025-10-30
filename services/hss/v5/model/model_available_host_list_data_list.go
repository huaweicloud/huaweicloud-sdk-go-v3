package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AvailableHostListDataList struct {

	// **参数解释**： 主机ID **取值范围**： 字符长度1-64位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**： 服务器私有IP **取值范围**： 字符长度1-128位
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**: Agent ID **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	AgentId *string `json:"agent_id,omitempty"`

	// 冲突的端口
	ConflictPort *[]int32 `json:"conflict_port,omitempty"`

	// 操作系统类型，包含如下2种。 - Linux ：Linux。 - Windows ：Windows.
	OsType *string `json:"os_type,omitempty"`

	// 分组名称
	GroupName *string `json:"group_name,omitempty"`

	// 分组id
	GroupId *string `json:"group_id,omitempty"`
}

func (o AvailableHostListDataList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailableHostListDataList struct{}"
	}

	return strings.Join([]string{"AvailableHostListDataList", string(data)}, " ")
}
