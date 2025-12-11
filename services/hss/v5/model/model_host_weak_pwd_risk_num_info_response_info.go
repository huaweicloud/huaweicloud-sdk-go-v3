package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HostWeakPwdRiskNumInfoResponseInfo 主机弱口令风险信息
type HostWeakPwdRiskNumInfoResponseInfo struct {

	// **参数解释**： 服务器（主机）的唯一标识ID **取值范围**： 字符长度1-64位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 服务器IP **取值范围**: 字符长度1-128位
	HostIp *string `json:"host_ip,omitempty"`

	// **参数解释**: 弱口令数量。 **取值范围**: 取值0-2147483647
	WeakPwdNum *int32 `json:"weak_pwd_num,omitempty"`
}

func (o HostWeakPwdRiskNumInfoResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostWeakPwdRiskNumInfoResponseInfo struct{}"
	}

	return strings.Join([]string{"HostWeakPwdRiskNumInfoResponseInfo", string(data)}, " ")
}
