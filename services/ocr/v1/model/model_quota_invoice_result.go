package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuotaInvoiceResult struct {
	// 发票号码。

	Number *string `json:"number,omitempty"`
	// 发票代码。

	Code *string `json:"code,omitempty"`
	// 地址。

	Location *string `json:"location,omitempty"`
	// 发票金额。

	Amount *string `json:"amount,omitempty"`
	// 相关字段的置信度信息，置信度越大，表示本次识别的对应字段的可靠性越高，在统计意义上，置信度越大，准确率越高。 置信度由算法给出，不直接等价于对应字段的准确率。

	Confidence *interface{} `json:"confidence,omitempty"`
}

func (o QuotaInvoiceResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaInvoiceResult struct{}"
	}

	return strings.Join([]string{"QuotaInvoiceResult", string(data)}, " ")
}
