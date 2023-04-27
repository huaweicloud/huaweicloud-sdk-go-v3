package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TenantBasic struct {

	// 租户custom id
	CustomerId *string `json:"customer_id,omitempty"`

	// 租户custom name
	CustomerName *string `json:"customer_name,omitempty"`

	// 租户企业名称
	EnterpriseName *string `json:"enterprise_name,omitempty"`
}

func (o TenantBasic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TenantBasic struct{}"
	}

	return strings.Join([]string{"TenantBasic", string(data)}, " ")
}
