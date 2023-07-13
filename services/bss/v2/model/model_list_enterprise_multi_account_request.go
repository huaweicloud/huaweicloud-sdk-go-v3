package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnterpriseMultiAccountRequest Request Object
type ListEnterpriseMultiAccountRequest struct {

	// 企业子账户的账号ID。
	SubCustomerId string `json:"sub_customer_id"`

	// 账户类型：BALANCE_TYPE_DEBIT：余额账户（默认）BALANCE_TYPE_CREDIT：信用账户
	BalanceType string `json:"balance_type"`

	// 偏移量，默认值为0。只有信用账户有效。
	Offset *int32 `json:"offset,omitempty"`

	// 每次查询条数，默认值为10。只有信用账户有效。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListEnterpriseMultiAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnterpriseMultiAccountRequest struct{}"
	}

	return strings.Join([]string{"ListEnterpriseMultiAccountRequest", string(data)}, " ")
}
