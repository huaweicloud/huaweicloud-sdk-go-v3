package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProtectStatisticsResponse Response Object
type ShowProtectStatisticsResponse struct {

	// **参数解释**: 漏洞库更新时间 **取值范围**: 最小值0，最大值4071095999000
	VulLibraryUpdateTime *int64 `json:"vul_library_update_time,omitempty"`

	// **参数解释**: 守护天数 **取值范围**: 最小值0，最大值4071095999000
	ProtectDays *int64 `json:"protect_days,omitempty"`

	// **参数解释**: 病毒库更新时间 **取值范围**: 最小值0，最大值4071095999000
	ThreatLibraryUpdateTime *int64 `json:"threat_library_update_time,omitempty"`

	// **参数解释**: 漏洞累计已检测数量 **取值范围**: 最小值0，最大值4071095999000
	VulDetectedTotalNum *int64 `json:"vul_detected_total_num,omitempty"`

	// **参数解释**: 累计检测基线总数量 **取值范围**: 最小值0，最大值4071095999000
	BaselineDetectedTotalNum *int64 `json:"baseline_detected_total_num,omitempty"`

	// **参数解释**: 累计扫描指纹数量 **取值范围**: 最小值0，最大值4071095999000
	FingerScanTotalNum *int64 `json:"finger_scan_total_num,omitempty"`

	// **参数解释**: 入侵检测累计检测告警总数量 **取值范围**: 最小值0，最大值4071095999000
	AlarmDetectedTotalNum *int64 `json:"alarm_detected_total_num,omitempty"`

	// **参数解释**: 累计防御勒索病毒告警次数 **取值范围**: 最小值0，最大值4071095999000
	RansomwareAlarmDetectedTotalNum *int64 `json:"ransomware_alarm_detected_total_num,omitempty"`

	// **参数解释**: 文件完整性监控累计检测文件变更告警总数 **取值范围**: 最小值0，最大值4071095999000
	FileAlarmDetectedTotalNum *int64 `json:"file_alarm_detected_total_num,omitempty"`

	// **参数解释**: 应用防护累计检测告警总数 **取值范围**: 最小值0，最大值4071095999000
	RaspAlarmDetectedTotalNum *int64 `json:"rasp_alarm_detected_total_num,omitempty"`

	// **参数解释**: 网页防篡改累计抵御网页防篡改次数 **取值范围**: 最小值0，最大值4071095999000
	WtpAlarmDetectedTotalNum *int64 `json:"wtp_alarm_detected_total_num,omitempty"`

	// **参数解释**: 容器镜像安全累计检测风险个数 **取值范围**: 最小值0，最大值4071095999000
	ImageRiskTotalNum *int64 `json:"image_risk_total_num,omitempty"`

	// **参数解释**: 容器安全防护累计检测容器告警个数 **取值范围**: 最小值0，最大值4071095999000
	ContainerAlarmTotalNum *int64 `json:"container_alarm_total_num,omitempty"`

	// **参数解释**: 容器防火墙累计设置策略条数 **取值范围**: 最小值0，最大值4071095999000
	ContainerFirewallPolicyTotalNum *int64 `json:"container_firewall_policy_total_num,omitempty"`

	// **参数解释**: 是否开启恶意自动查杀 **取值范围**: - true：是。 - false：否。
	AutoKillVirusStatus *bool `json:"auto_kill_virus_status,omitempty"`
	HttpStatusCode      int   `json:"-"`
}

func (o ShowProtectStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProtectStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowProtectStatisticsResponse", string(data)}, " ")
}
