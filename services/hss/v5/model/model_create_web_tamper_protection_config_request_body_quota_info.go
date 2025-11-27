package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWebTamperProtectionConfigRequestBodyQuotaInfo **参数解释**: 开启防护时绑定的配额信息 **约束限制**: open_protection_immediately值为true时有效 **取值范围**: 不涉及 **默认取值**: 不涉及
type CreateWebTamperProtectionConfigRequestBodyQuotaInfo struct {

	// **参数解释**: 计费模式 **约束限制**: 不涉及 **取值范围**: - packet_cycle：包年包月  **默认取值**: 不涉及
	ChargingMode string `json:"charging_mode"`

	// **参数解释**: 配额id列表 **约束限制**: 若该字段值为空，则随机选取配额 **取值范围**: 最少0条，最多10条 **默认取值**: 不涉及
	ResourceIdList *[]string `json:"resource_id_list,omitempty"`
}

func (o CreateWebTamperProtectionConfigRequestBodyQuotaInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWebTamperProtectionConfigRequestBodyQuotaInfo struct{}"
	}

	return strings.Join([]string{"CreateWebTamperProtectionConfigRequestBodyQuotaInfo", string(data)}, " ")
}
