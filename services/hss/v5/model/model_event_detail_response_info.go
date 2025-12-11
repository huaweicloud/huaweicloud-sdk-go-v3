package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventDetailResponseInfo 操作详情信息
type EventDetailResponseInfo struct {

	// **参数解释**: 主机上安装的杀毒Agent的唯一标识ID，用于关联主机与杀毒服务 **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释**： 进程ID **取值范围**： 最小值0，最大值2147483647
	ProcessPid *int32 `json:"process_pid,omitempty"`

	// **参数解释**： 是否是父进程 **取值范围**： - true：是父进程 - false：不是父进程
	IsParent *bool `json:"is_parent,omitempty"`

	// **参数解释**： 文件哈希 **取值范围**： 字符长度1-256位
	FileHash *string `json:"file_hash,omitempty"`

	// **参数解释**： 文件路径 **取值范围**： 字符长度1-256位
	FilePath *string `json:"file_path,omitempty"`

	// **参数解释**： 文件的系统属性（如读写权限、隐藏属性、执行权限等） **取值范围**： 字符长度1-256位
	FileAttr *string `json:"file_attr,omitempty"`

	// **参数解释**： 服务器私有IP **取值范围**： 字符长度1-128位
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**： 登录源IP **取值范围**： 字符长度1-256位
	LoginIp *string `json:"login_ip,omitempty"`

	// **参数解释**： 登录用户名 **取值范围**： 字符长度1-256位
	LoginUserName *string `json:"login_user_name,omitempty"`

	// 告警事件关键字，仅用于告警白名单
	Keyword *string `json:"keyword,omitempty"`

	// 告警事件hash，仅用于告警白名单
	Hash *string `json:"hash,omitempty"`
}

func (o EventDetailResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventDetailResponseInfo struct{}"
	}

	return strings.Join([]string{"EventDetailResponseInfo", string(data)}, " ")
}
