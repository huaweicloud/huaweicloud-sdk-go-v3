package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityCheckHostInfoResponseInfo 受配置检测影响的服务器信息
type SecurityCheckHostInfoResponseInfo struct {

	// **参数解释**: 主机ID **取值范围**: 不涉及
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 服务器名称 **取值范围**: 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 服务器公网IP **取值范围**: 不涉及
	HostPublicIp *string `json:"host_public_ip,omitempty"`

	// **参数解释**: 服务器私网IP **取值范围**: 不涉及
	HostPrivateIp *string `json:"host_private_ip,omitempty"`

	// **参数解释**: 扫描时间(ms) **取值范围**: 不涉及
	ScanTime *int64 `json:"scan_time,omitempty"`

	// **参数解释**: 风险项数量 **取值范围**: 不涉及
	FailedNum *int32 `json:"failed_num,omitempty"`

	// **参数解释**: 通过项数量 **取值范围**: 不涉及
	PassedNum *int32 `json:"passed_num,omitempty"`
}

func (o SecurityCheckHostInfoResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityCheckHostInfoResponseInfo struct{}"
	}

	return strings.Join([]string{"SecurityCheckHostInfoResponseInfo", string(data)}, " ")
}
