package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTenantQuotasResponseBodyQuotas 配额列表对象。
type ListTenantQuotasResponseBodyQuotas struct {

	// 配额资源列表。
	Resources *[]QuotaOuterResource `json:"resources,omitempty"`
}

func (o ListTenantQuotasResponseBodyQuotas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTenantQuotasResponseBodyQuotas struct{}"
	}

	return strings.Join([]string{"ListTenantQuotasResponseBodyQuotas", string(data)}, " ")
}
