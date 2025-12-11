package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventDetailRequestInfo 操作详情信息
type EventDetailRequestInfo struct {

	// **参数解释**: 主机上安装的杀毒Agent的唯一标识ID，用于关联主机与杀毒服务 **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释**： 进程ID **取值范围**： 最小值0，最大值2147483647
	ProcessPid *int32 `json:"process_pid,omitempty"`

	// **参数解释**： 文件哈希 **取值范围**： 字符长度1-256位
	FileHash *string `json:"file_hash,omitempty"`

	// **参数解释**： 文件路径 **取值范围**： 字符长度1-256位
	FilePath *string `json:"file_path,omitempty"`

	// **参数解释**： 文件的系统属性（如读写权限、隐藏属性、执行权限等） **取值范围**： 字符长度1-256位
	FileAttr *string `json:"file_attr,omitempty"`

	// **参数解释**： 告警事件关键字，仅用于告警白名单 **取值范围**： 字符长度1-256位
	Keyword *string `json:"keyword,omitempty"`

	// **参数解释**： 告警事件hash，仅用于告警白名单 **取值范围**： 字符长度1-256位
	Hash *string `json:"hash,omitempty"`

	// **参数解释**： 服务器私有IP **取值范围**： 字符长度1-128位
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**： 登录源IP **取值范围**： 字符长度1-256位
	LoginIp *string `json:"login_ip,omitempty"`

	// **参数解释**： 登录用户名 **取值范围**： 字符长度1-256位
	LoginUserName *string `json:"login_user_name,omitempty"`

	// 容器ID
	ContainerId *string `json:"container_id,omitempty"`

	// 容器名称
	ContainerName *string `json:"container_name,omitempty"`
}

func (o EventDetailRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventDetailRequestInfo struct{}"
	}

	return strings.Join([]string{"EventDetailRequestInfo", string(data)}, " ")
}
