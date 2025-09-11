package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTenantQuotasResponse Response Object
type ListTenantQuotasResponse struct {
	Quotas         *ListTenantQuotasResponseBodyQuotas `json:"quotas,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ListTenantQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTenantQuotasResponse struct{}"
	}

	return strings.Join([]string{"ListTenantQuotasResponse", string(data)}, " ")
}
