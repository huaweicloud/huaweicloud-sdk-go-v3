package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryTenantQuotaRespQuotas 配额信息。
type QueryTenantQuotaRespQuotas struct {

	// 配额列表。
	Resources *[]Resources `json:"resources,omitempty"`
}

func (o QueryTenantQuotaRespQuotas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryTenantQuotaRespQuotas struct{}"
	}

	return strings.Join([]string{"QueryTenantQuotaRespQuotas", string(data)}, " ")
}
