package model

import (
	"encoding/json"

	"strings"
)

//
type BankcardResult struct {
	// 发卡行。

	BankName *string `json:"bank_name,omitempty"`
	// 银行卡号。

	CardNumber *string `json:"card_number,omitempty"`
	// 有效期开始日期。

	IssueDate *string `json:"issue_date,omitempty"`
	// 有效期截止日期。

	ExpiryDate *string `json:"expiry_date,omitempty"`
	// 银行卡类别，如：储蓄卡，信用卡。

	Type *string `json:"type,omitempty"`
	// 相关字段的置信度信息，置信度越大，表示本次识别的对应字段的可靠性越高，在统计意义上，置信度越大，准确率越高。 置信度由算法给出，不直接等价于对应字段的准确率。

	Confidence *interface{} `json:"confidence,omitempty"`
}

func (o BankcardResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BankcardResult struct{}"
	}

	return strings.Join([]string{"BankcardResult", string(data)}, " ")
}
