package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WeakPwdListInfoResponseInfo 服务器弱口令列表
type WeakPwdListInfoResponseInfo struct {

	// **参数解释**: 主机ID **取值范围**: 不涉及
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 服务器名称 **取值范围**: 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 服务器IP（私有IP），为兼容用户使用，不删除此字段 **取值范围**: 不涉及
	HostIp *string `json:"host_ip,omitempty"`

	// **参数解释**: 服务器私有IP **取值范围**: 不涉及
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**: 服务器公网IP **取值范围**: 不涉及
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**: 最近扫描时间，时间戳单位：毫秒 **取值范围**: 不涉及
	UpdateTime *int64 `json:"update_time,omitempty"`

	// **参数解释**: 弱口令账号列表 **取值范围**: 不涉及
	WeakPwdAccounts *[]WeakPwdAccountInfoResponseInfo `json:"weak_pwd_accounts,omitempty"`
}

func (o WeakPwdListInfoResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WeakPwdListInfoResponseInfo struct{}"
	}

	return strings.Join([]string{"WeakPwdListInfoResponseInfo", string(data)}, " ")
}
