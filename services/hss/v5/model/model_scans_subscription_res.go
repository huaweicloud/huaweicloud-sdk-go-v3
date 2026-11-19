package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScansSubscriptionRes 风险预防订阅的项目
type ScansSubscriptionRes struct {

	// **参数解释**: 漏洞统计 **取值范围**: - true：开启 - false：关闭
	VulnerabilitiesStatisticsStatus *bool `json:"vulnerabilities_statistics_status,omitempty"`

	// **参数解释**: 资产账号变动记录 **取值范围**: - true：开启 - false：关闭
	AccountChangeHistoryStatus *bool `json:"account_change_history_status,omitempty"`

	// **参数解释**: 危险开放端口 **取值范围**: - true：开启 - false：关闭
	DangerPortsStatus *bool `json:"danger_ports_status,omitempty"`

	// **参数解释**: 弱口令 **取值范围**: - true：开启 - false：关闭
	WeakPwdStatus *bool `json:"weak_pwd_status,omitempty"`
}

func (o ScansSubscriptionRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScansSubscriptionRes struct{}"
	}

	return strings.Join([]string{"ScansSubscriptionRes", string(data)}, " ")
}
