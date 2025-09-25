package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProtectInfoQuotaInfo 防护配额信息
type ProtectInfoQuotaInfo struct {

	// **参数解释**: 已经到期的配额 **取值范围**: 最小值0，最大值2147483647
	ExpiredQuotaNum *int32 `json:"expired_quota_num,omitempty"`

	// **参数解释**: 即将到期的配额 **取值范围**: 最小值0，最大值2147483647
	AboutToExpireQuotaNum *int32 `json:"about_to_expire_quota_num,omitempty"`

	// **参数解释**: 使用企业版的用户率 **取值范围**: 最小值0，最大值1
	UserUseEnterpriseRate *float32 `json:"user_use_enterprise_rate,omitempty"`
}

func (o ProtectInfoQuotaInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectInfoQuotaInfo struct{}"
	}

	return strings.Join([]string{"ProtectInfoQuotaInfo", string(data)}, " ")
}
