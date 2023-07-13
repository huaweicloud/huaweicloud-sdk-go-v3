package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMultiAccountTransferAmountRequest Request Object
type ShowMultiAccountTransferAmountRequest struct {

	// 账户类型：BALANCE_TYPE_DEBIT：余额账户BALANCE_TYPE_CREDIT：信用账户
	BalanceType string `json:"balance_type"`

	// 偏移量，默认值为0。只有信用账户有效。
	Offset *int32 `json:"offset,omitempty"`

	// 每次查询条数，默认值为10。只有信用账户有效。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowMultiAccountTransferAmountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMultiAccountTransferAmountRequest struct{}"
	}

	return strings.Join([]string{"ShowMultiAccountTransferAmountRequest", string(data)}, " ")
}
