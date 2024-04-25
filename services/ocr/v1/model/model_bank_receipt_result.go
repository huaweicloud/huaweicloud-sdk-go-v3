package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BankReceiptResult struct {

	// 银行回单数量
	BankReceiptCount *int32 `json:"bank_receipt_count,omitempty"`

	// 银行回单键值对提取结果。
	BankReceiptList *[]BankReceiptDict `json:"bank_receipt_list,omitempty"`
}

func (o BankReceiptResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BankReceiptResult struct{}"
	}

	return strings.Join([]string{"BankReceiptResult", string(data)}, " ")
}
