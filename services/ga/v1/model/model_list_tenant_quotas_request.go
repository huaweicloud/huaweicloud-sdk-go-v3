package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTenantQuotasRequest Request Object
type ListTenantQuotasRequest struct {

	// 租户ID。
	DomainId string `json:"domain_id"`
}

func (o ListTenantQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTenantQuotasRequest struct{}"
	}

	return strings.Join([]string{"ListTenantQuotasRequest", string(data)}, " ")
}
