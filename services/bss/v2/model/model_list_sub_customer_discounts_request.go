package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListSubCustomerDiscountsRequest struct {
	// 客户账号ID。您可以调用查询客户列表接口获取customer_id。

	CustomerId string `json:"customer_id"`
	// 精英服务商ID。获取方法请参见查询精英服务商列表。

	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
}

func (o ListSubCustomerDiscountsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubCustomerDiscountsRequest struct{}"
	}

	return strings.Join([]string{"ListSubCustomerDiscountsRequest", string(data)}, " ")
}
