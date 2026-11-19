package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProtectionStatusStatisticsResponse Response Object
type ShowProtectionStatusStatisticsResponse struct {

	// **参数解释**: 漏洞扫描时间，Unix时间戳格式，单位为毫秒 **取值范围**: 最小值0，最大值9223372036854775807
	VulScanTime *int64 `json:"vul_scan_time,omitempty"`

	// **参数解释**: 病毒库更新时间，Unix时间戳格式，单位为毫秒 **取值范围**: 最小值0，最大值9223372036854775807
	AvUpgradeTime *int64 `json:"av_upgrade_time,omitempty"`

	// **参数解释**: 已开启防护的类型列表 **取值范围**: 0-200个ProtectionEnabledTypes对象
	ProtectionEnabledList *[]ProtectionEnabledTypes `json:"protection_enabled_list,omitempty"`

	// **参数解释**: 防护状态统计列表 **取值范围**: 0-200个ProtectStatusStatistics对象
	DataList       *[]ProtectStatusStatistics `json:"data_list,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ShowProtectionStatusStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProtectionStatusStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowProtectionStatusStatisticsResponse", string(data)}, " ")
}
