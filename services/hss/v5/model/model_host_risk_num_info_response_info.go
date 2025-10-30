package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HostRiskNumInfoResponseInfo 服务器风险数量信息
type HostRiskNumInfoResponseInfo struct {

	// **参数解释**： 主机ID **取值范围**： 字符长度1-64位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 服务器IP **取值范围**: 字符长度1-128位
	HostIp *string `json:"host_ip,omitempty"`

	// **参数解释**: 扫描时间(ms)。 **取值范围**: 取值0-9223372036854775807
	ScanTime *int64 `json:"scan_time,omitempty"`

	// **参数解释**: 高危风险数量。 **取值范围**: 取值0-2147483647
	HighRiskNum *int32 `json:"high_risk_num,omitempty"`

	// **参数解释**: 中危风险数量。 **取值范围**: 取值0-2147483647
	MediumRiskNum *int32 `json:"medium_risk_num,omitempty"`

	// **参数解释**: 低危风险数量。 **取值范围**: 取值0-2147483647
	LowRiskNum *int32 `json:"low_risk_num,omitempty"`
}

func (o HostRiskNumInfoResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostRiskNumInfoResponseInfo struct{}"
	}

	return strings.Join([]string{"HostRiskNumInfoResponseInfo", string(data)}, " ")
}
