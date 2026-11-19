package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScansSubscriptionReq 扫描订阅信息
type ScansSubscriptionReq struct {

	// **参数解释**： 漏洞统计是否订阅 **约束限制**： 不涉及 **取值范围**： - true：是 - false：否  **默认取值**： 不涉及
	VulnerabilitiesStatisticsStatus *bool `json:"vulnerabilities_statistics_status,omitempty"`

	// **参数解释**： 资产账号变动记录是否订阅 **约束限制**： 不涉及 **取值范围**： - true：是 - false：否  **默认取值**： 不涉及
	AccountChangeHistoryStatus *bool `json:"account_change_history_status,omitempty"`

	// **参数解释**： 危险开放端口是否订阅 **约束限制**： 不涉及 **取值范围**： - true：是 - false：否  **默认取值**： 不涉及
	DangerPortsStatus *bool `json:"danger_ports_status,omitempty"`

	// **参数解释**： 弱口令是否订阅 **约束限制**： 不涉及 **取值范围**： - true：是 - false：否  **默认取值**： 不涉及
	WeakPwdStatus *bool `json:"weak_pwd_status,omitempty"`
}

func (o ScansSubscriptionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScansSubscriptionReq struct{}"
	}

	return strings.Join([]string{"ScansSubscriptionReq", string(data)}, " ")
}
