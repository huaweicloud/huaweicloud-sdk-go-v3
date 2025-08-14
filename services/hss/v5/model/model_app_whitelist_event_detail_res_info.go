package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AppWhitelistEventDetailResInfo 操作详情信息
type AppWhitelistEventDetailResInfo struct {

	// **参数解释**: Agent ID **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释**： 进程ID **取值范围**： 最小值0，最大值2147483647
	ProcessPid *int32 `json:"process_pid,omitempty"`

	// **参数解释**： 文件哈希 **取值范围**： 字符长度1-256位
	FileHash *string `json:"file_hash,omitempty"`

	// **参数解释**： 文件路径 **取值范围**： 字符长度1-256位
	FilePath *string `json:"file_path,omitempty"`

	// **参数解释**： 文件属性 **取值范围**： 字符长度1-256位
	FileAttr *string `json:"file_attr,omitempty"`

	// **参数解释**： 服务器私有IP **取值范围**： 字符长度1-128位
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**： 登录源IP **取值范围**： 字符长度1-256位
	LoginIp *string `json:"login_ip,omitempty"`

	// **参数解释**： 登录用户名 **取值范围**： 字符长度1-256位
	LoginUserName *string `json:"login_user_name,omitempty"`
}

func (o AppWhitelistEventDetailResInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppWhitelistEventDetailResInfo struct{}"
	}

	return strings.Join([]string{"AppWhitelistEventDetailResInfo", string(data)}, " ")
}
