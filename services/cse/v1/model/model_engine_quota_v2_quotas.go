package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EngineQuotaV2Quotas struct {
	Resources *[]TenantQuotaUsed `json:"resources,omitempty"`
}

func (o EngineQuotaV2Quotas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EngineQuotaV2Quotas struct{}"
	}

	return strings.Join([]string{"EngineQuotaV2Quotas", string(data)}, " ")
}
