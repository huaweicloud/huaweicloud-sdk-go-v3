package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventDetailResponseInfo 操作详情信息
type EventDetailResponseInfo struct {

	// Agent ID
	AgentId *string `json:"agent_id,omitempty"`

	// 进程id
	ProcessPid *int32 `json:"process_pid,omitempty"`

	// 是否是父进程
	IsParent *bool `json:"is_parent,omitempty"`

	// 文件哈希
	FileHash *string `json:"file_hash,omitempty"`

	// 文件路径
	FilePath *string `json:"file_path,omitempty"`

	// 文件属性
	FileAttr *string `json:"file_attr,omitempty"`

	// 服务器私有IP
	PrivateIp *string `json:"private_ip,omitempty"`

	// 登录源IP
	LoginIp *string `json:"login_ip,omitempty"`

	// 登录用户名
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
