package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecurityCheckHostReportResponse Response Object
type ShowSecurityCheckHostReportResponse struct {
	HostResult *SecurityCheckHostResultResponseInfo `json:"host_result,omitempty"`

	// **参数解释**： 服务器运行状态 **取值范围**： - ACTIVE：运行中 - SHUTOFF：关机 - BUILDING：创建中 - ERROR：故障
	HostStatus *string `json:"host_status,omitempty"`

	// **参数解释**： 体检周期开始时间 **取值范围**： 不涉及
	ScanPeriodStart *int64 `json:"scan_period_start,omitempty"`

	// **参数解释**： 体检周期结束时间 **取值范围**： 不涉及
	ScanPeriodEnd *int64 `json:"scan_period_end,omitempty"`

	// **参数解释**： 风险评分 **取值范围**： 不涉及
	RiskRating *int32 `json:"risk_rating,omitempty"`

	// **参数解释**： 风险级别 **取值范围**： - high：高危 - medium：中危 - low：低危 - safe：安全，无风险
	Severity *string `json:"severity,omitempty"`

	// **参数解释**： 安全风险数量 **取值范围**： 不涉及
	RiskNum *int32 `json:"risk_num,omitempty"`

	// **参数解释**： 最新检测时间 **取值范围**： 不涉及
	ScanTime *int64 `json:"scan_time,omitempty"`

	// **参数解释**: 是否是免费安全体检的报告 **取值范围**: - true：免费安全体检的报告 - false：非免费安全体检的报告
	Free *bool `json:"free,omitempty"`

	EventNumInfo *SecurityCheckRiskNumInfo `json:"event_num_info,omitempty"`

	VulNumInfo *SecurityCheckRiskNumInfo `json:"vul_num_info,omitempty"`

	BaselineNumInfo *SecurityCheckRiskNumInfo `json:"baseline_num_info,omitempty"`

	AssetNumInfo *SecurityCheckRiskNumInfo `json:"asset_num_info,omitempty"`

	// **参数解释**： 主机资产变动记录数量 **取值范围**： 不涉及
	AssetChangeNum *int32 `json:"asset_change_num,omitempty"`

	// **参数解释**： 软件数量 **取值范围**： 不涉及
	AppNum *int32 `json:"app_num,omitempty"`

	// **参数解释**： 危险端口数量 **取值范围**： 不涉及
	PortNum *int32 `json:"port_num,omitempty"`

	// **参数解释**： 入侵事件列表 **取值范围**： 不涉及
	EventList *[]SecurityCheckRiskEventInfo `json:"event_list,omitempty"`

	// **参数解释**： 漏洞列表 **取值范围**： 不涉及
	VulList *[]SecurityCheckVulInfo `json:"vul_list,omitempty"`

	// **参数解释**： 配置检测列表 **取值范围**： 不涉及
	SecurityConfigList *[]SecurityConfigInfo `json:"security_config_list,omitempty"`

	// **参数解释**： 配置检查项列表 **取值范围**： 不涉及
	SecurityConfigItemList *[]SecurityConfigItemInfo `json:"security_config_item_list,omitempty"`

	// **参数解释**： 口令复杂度策略列表 **取值范围**： 不涉及
	PwdPolicyList *[]SecurityConfigPwdPolicyInfo `json:"pwd_policy_list,omitempty"`

	// **参数解释**： 经典弱口令列表 **取值范围**： 不涉及
	WeakPwdList *[]SecurityConfigWeakPwdInfo `json:"weak_pwd_list,omitempty"`

	// **参数解释**： 主机账户的历史变动记录 **取值范围**： 不涉及
	UserChangeList *[]SecurityConfigUserChangeInfo `json:"user_change_list,omitempty"`

	// **参数解释**： 危险开放端口列表 **取值范围**： 不涉及
	PortList *[]SecurityConfigPortInfo `json:"port_list,omitempty"`

	// **参数解释**： 资产的软件列表 **取值范围**： 不涉及
	AppList        *[]AppResponseInfo `json:"app_list,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowSecurityCheckHostReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityCheckHostReportResponse struct{}"
	}

	return strings.Join([]string{"ShowSecurityCheckHostReportResponse", string(data)}, " ")
}
