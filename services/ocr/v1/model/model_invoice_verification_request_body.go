package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InvoiceVerificationRequestBody
type InvoiceVerificationRequestBody struct {

	// 发票代码。发票种类为全电发票时，该参数须为空字符串。
	Code string `json:"code"`

	// 发票号码
	Number string `json:"number"`

	// 发票日期格式YYYY-MM-DD
	IssueDate string `json:"issue_date"`

	// 校验码后六位。 - 以下种类发票，参数不可为空    增值税普通发票、增值税电子普通发票、增值税普通发票（卷式）、增值税电子普通发票（通行费）、区块链电子发票。  - 区块链电子发票需要填写5位校验码。
	CheckCode *string `json:"check_code,omitempty"`

	// 合计金额。和票据上的金额的有效数字保持一致，例如票据上的金额为88.00，则需要输入字符串为“88.00”，才能验真成功。如果输入“88”或“88.0”可能会产生\"result_code\": \"1010\", \" Parameter error.\"报错。  - 以下种类发票，参数不可为空    增值税专用发票、增值税电子专用发票、机动车销售统一发票、二手车销售统一发票、区块链电子发票、全电发票。  - 填写发票合计金额（不含税）    增值税专用发票、增值税电子专用发票、机动车销售统一发票、区块链电子发票。  - 填写发票车价合计    二手车发票。  - 填写发票合计金额    全电发票。  - 填写票价或者退票费    全电发票（铁路电子客票）
	SubtotalAmount *string `json:"subtotal_amount,omitempty"`
}

func (o InvoiceVerificationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InvoiceVerificationRequestBody struct{}"
	}

	return strings.Join([]string{"InvoiceVerificationRequestBody", string(data)}, " ")
}
