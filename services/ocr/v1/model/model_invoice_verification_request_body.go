package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type InvoiceVerificationRequestBody struct {

	// 发票代码
	Code string `json:"code"`

	// 发票号码
	Number string `json:"number"`

	// 发票日期格式YYYY-MM-DD
	IssueDate string `json:"issue_date"`

	// 校验码后六位  发票种类为增值税普通发票、增值税电子普通发票、增值税普通发票（卷式）、增值税电子普通发票（通行费）、区块链电子发票时此项不可为空（区块链电子发票验真时，填写的是5位校验码）
	CheckCode *string `json:"check_code,omitempty"`

	// 合计金额（不含税）  1.发票种类为增值税专用发票、增值税电子专用发票、机动车销售统一发票、二手车销售统一发票、区块链电子发票时不可为空； 2.增值税专用发票、增值税电子专用发票、机动车销售统一发票、区块链电子发票填写发票合计金额（不含税），二手车发票填写发票车价合计
	SubtotalAmount *string `json:"subtotal_amount,omitempty"`
}

func (o InvoiceVerificationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InvoiceVerificationRequestBody struct{}"
	}

	return strings.Join([]string{"InvoiceVerificationRequestBody", string(data)}, " ")
}
