package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 操作详情信息
type EventDetailResponseInfo struct {

	// Agent ID
	AgentId *string `json:"agent_id,omitempty" xml:"agent_id"`

	// 进程id
	ProcessPid *int32 `json:"process_pid,omitempty" xml:"process_pid"`

	// 是否是父进程
	IsParent *bool `json:"is_parent,omitempty" xml:"is_parent"`

	// 文件哈希
	FileHash *string `json:"file_hash,omitempty" xml:"file_hash"`

	// 文件路径
	FilePath *string `json:"file_path,omitempty" xml:"file_path"`

	// 文件属性
	FileAttr *string `json:"file_attr,omitempty" xml:"file_attr"`

	// 服务器私网IP
	PrivateIp *string `json:"private_ip,omitempty" xml:"private_ip"`

	// 登录源IP
	LoginIp *string `json:"login_ip,omitempty" xml:"login_ip"`

	// 登录用户名
	LoginUserName *string `json:"login_user_name,omitempty" xml:"login_user_name"`
}

func (o EventDetailResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventDetailResponseInfo struct{}"
	}

	return strings.Join([]string{"EventDetailResponseInfo", string(data)}, " ")
}
