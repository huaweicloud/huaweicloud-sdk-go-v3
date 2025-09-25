package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RiskHandleInfoAlarmInfo 告警信息
type RiskHandleInfoAlarmInfo struct {

	// **参数解释**: 攻击、异常行为数 **取值范围**: 最小值0，最大值2147483647
	AlarmNum *int64 `json:"alarm_num,omitempty"`

	// **参数解释**: 病毒文件数 **取值范围**: 最小值0，最大值2147483647
	MalwareFileNum *int64 `json:"malware_file_num,omitempty"`

	// **参数解释**: 阻断的数量 **取值范围**: 最小值0，最大值2147483647
	AutoBlockNum *int64 `json:"auto_block_num,omitempty"`

	// **参数解释**: 手工处置数量 **取值范围**: 最小值0，最大值2147483647
	ManualHandleNum *int64 `json:"manual_handle_num,omitempty"`

	// **参数解释**: 自动处置数量 **取值范围**: 最小值0，最大值2147483647
	AutoHandleNum *int64 `json:"auto_handle_num,omitempty"`

	// **参数解释**: 平均处置率 **取值范围**: 最小值0，最大值1
	HandleRate *float32 `json:"handle_rate,omitempty"`

	// **参数解释**: 平均处置率击败的用户率 **取值范围**: 最小值0，最大值1
	BeatRate *float32 `json:"beat_rate,omitempty"`

	// **参数解释**: 是否开启隔离查杀 **取值范围**:  - true：是。  - false：否。
	VirusKillConfigEnable *bool `json:"virus_kill_config_enable,omitempty"`

	// **参数解释**: 开启隔离查杀的用户率 **取值范围**: 最小值0，最大值1
	UserOpenVirusKillRate *float32 `json:"user_open_virus_kill_rate,omitempty"`

	// **参数解释**: 是否开启告警通知 **取值范围**:  - true：是。  - false：否。
	AlarmNotifyConfigEnable *bool `json:"alarm_notify_config_enable,omitempty"`

	// **参数解释**: 开启告警通知的用户率 **取值范围**: 最小值0，最大值1
	UserOpenAlarmNotifyRate *float32 `json:"user_open_alarm_notify_rate,omitempty"`

	// **参数解释**: 勒索病毒攻击数量 **取值范围**: 最小值0，最大值2147483647
	RansomwareEventNum *int64 `json:"ransomware_event_num,omitempty"`

	// **参数解释**: 存在勒索病毒攻击的主机列表 **取值范围**: 最小值0，最大值10000
	RansomwareEventHostList *[]string `json:"ransomware_event_host_list,omitempty"`
}

func (o RiskHandleInfoAlarmInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RiskHandleInfoAlarmInfo struct{}"
	}

	return strings.Join([]string{"RiskHandleInfoAlarmInfo", string(data)}, " ")
}
