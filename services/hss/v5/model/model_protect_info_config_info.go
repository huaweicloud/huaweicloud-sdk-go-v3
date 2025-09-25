package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProtectInfoConfigInfo 防护配置信息
type ProtectInfoConfigInfo struct {

	// **参数解释**: 未升级agent的主机数 **取值范围**: 最小值0，最大值2147483647
	UnUpgradeAgentHostNum *int32 `json:"un_upgrade_agent_host_num,omitempty"`

	// **参数解释**: 是否开启自动升级 **取值范围**:   - true：是。   - false：否。
	AutoUpgradeAgentEnable *bool `json:"auto_upgrade_agent_enable,omitempty"`

	// **参数解释**: 开启自动升级的用户率 **取值范围**: 最小值0，最大值1
	UserOpenAutoUpgradeAgentRate *float32 `json:"user_open_auto_upgrade_agent_rate,omitempty"`

	// **参数解释**: 未开启勒索策略的主机数 **取值范围**: 最小值0，最大值2147483647
	UnOpenRansomwarePolicyHostNum *int32 `json:"un_open_ransomware_policy_host_num,omitempty"`

	// **参数解释**: 未待开启（待优化）勒索病毒防护的主机数 **取值范围**: 最小值0，最大值2147483647
	ToBeOptimizedRansomwarePolicyHostNum *int32 `json:"to_be_optimized_ransomware_policy_host_num,omitempty"`

	// **参数解释**: 开启勒索备份的用户率 **取值范围**: 最小值0，最大值1
	UserOpenRansomwareBackupRate *float32 `json:"user_open_ransomware_backup_rate,omitempty"`

	// **参数解释**: 未开启双因子的主机数 **取值范围**: 最小值0，最大值2147483647
	UnOpenTwoFactorLoginHostNum *int32 `json:"un_open_two_factor_login_host_num,omitempty"`

	// **参数解释**: 开启双因子的用户率 **取值范围**: 最小值0，最大值1
	UserOpenTwoFactorLoginRate *float32 `json:"user_open_two_factor_login_rate,omitempty"`

	// **参数解释**: 未执行病毒查杀的主机数 **取值范围**: 最小值0，最大值2147483647
	UnVirusKillHostNum *int32 `json:"un_virus_kill_host_num,omitempty"`

	// **参数解释**: 执行过病毒查杀的主机数 **取值范围**: 最小值0，最大值2147483647
	VirusKillHostNum *int32 `json:"virus_kill_host_num,omitempty"`

	// **参数解释**: 是否开启云查 **取值范围**:   - true：是。   - false：否。
	MalwareCollectEnable *bool `json:"malware_collect_enable,omitempty"`

	// **参数解释**: 开启云查的用户率 **取值范围**: 最小值0，最大值1
	UserOpenMalwareCollectRate *float32 `json:"user_open_malware_collect_rate,omitempty"`

	// **参数解释**: 容器镜像扫描频率 **取值范围**: 最小值0，最大值128
	ContainerImageScanFreq *float32 `json:"container_image_scan_freq,omitempty"`

	// **参数解释**: 容器镜像扫描频率击败的用户率 **取值范围**: 最小值0，最大值1
	ContainerImageScanFreqBeatRate *float32 `json:"container_image_scan_freq_beat_rate,omitempty"`

	// **参数解释**: 用户开启的配置数量 **取值范围**: 最小值0，最大值2147483647
	NeedOpenConfigNum *int32 `json:"need_open_config_num,omitempty"`

	// **参数解释**: 扫描的镜像数 **取值范围**: 最小值0，最大值2147483647
	ContainerImageScanTotalNum *int64 `json:"container_image_scan_total_num,omitempty"`
}

func (o ProtectInfoConfigInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectInfoConfigInfo struct{}"
	}

	return strings.Join([]string{"ProtectInfoConfigInfo", string(data)}, " ")
}
