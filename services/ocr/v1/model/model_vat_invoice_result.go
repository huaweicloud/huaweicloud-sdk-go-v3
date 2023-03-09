package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VatInvoiceResult struct {

	// 增值税发票标题
	Title *string `json:"title,omitempty"`

	// 增值税发票类型，可选值包括： - special: 增值税专用发票  - normal: 增值税普通发票  - electronic: 增值税电子普通发票  - special_electronic: 增值税电子专用发票  - toll: 增值税电子普通发票（通行费）  - roll: 增值税普通发票（卷票）  - fully_digitalized_special_electronic: 全电专用发票  - fully_digitalized_normal_electronic: 全电普通发票
	Type *string `json:"type,omitempty"`

	// 发票联次。 当“advanced_mode”设置为“true”时才返回。
	SerialNumber *string `json:"serial_number,omitempty"`

	// 发票归属地。 当“advanced_mode”设置为“true”时才返回。
	Attribution *string `json:"attribution,omitempty"`

	// 发票监制章。 当“advanced_mode”设置为“true”时才返回。
	SupervisionSeal *[]string `json:"supervision_seal,omitempty"`

	// 发票代码。
	Code *string `json:"code,omitempty"`

	// 机打代码。当“advanced_mode”设置为“true”时才返回。
	PrintCode *string `json:"print_code,omitempty"`

	// 机器编号。 当“advanced_mode”设置为“true”时才返回。
	MachineNumber *string `json:"machine_number,omitempty"`

	// 机打号码。 当“advanced_mode”设置为“true”时才返回
	PrintNumber *string `json:"print_number,omitempty"`

	// 发票校验码，特定类型增值税发票内不存在该信息时返回空字符串。
	CheckCode *string `json:"check_code,omitempty"`

	// 发票号码。
	Number *string `json:"number,omitempty"`

	// 开票日期。
	IssueDate *string `json:"issue_date,omitempty"`

	// 密码区。
	EncryptionBlock *string `json:"encryption_block,omitempty"`

	// 购买方名称。
	BuyerName *string `json:"buyer_name,omitempty"`

	// 购买方纳税人识别号。
	BuyerId *string `json:"buyer_id,omitempty"`

	// 购买方地址、电话。
	BuyerAddress *string `json:"buyer_address,omitempty"`

	// 购买方开户行及帐号。
	BuyerBank *string `json:"buyer_bank,omitempty"`

	// 销售方名称。
	SellerName *string `json:"seller_name,omitempty"`

	// 销售方纳税人识别号。
	SellerId *string `json:"seller_id,omitempty"`

	// 销售方地址、电话。
	SellerAddress *string `json:"seller_address,omitempty"`

	// 销售方开户行及帐号。
	SellerBank *string `json:"seller_bank,omitempty"`

	// 合计金额。
	SubtotalAmount *string `json:"subtotal_amount,omitempty"`

	// 合计税额。
	SubtotalTax *string `json:"subtotal_tax,omitempty"`

	// 价税合计。
	Total *string `json:"total,omitempty"`

	// 价税合计（大写）。 当“advanced_mode”设置为“true”时才返回。
	TotalInWords *string `json:"total_in_words,omitempty"`

	// 备注。 当“advanced_mode”设置为“true”时才返回。
	Remarks *string `json:"remarks,omitempty"`

	// 收款人。 当“advanced_mode”设置为“true”时才返回。
	Receiver *string `json:"receiver,omitempty"`

	// 复核。 当“advanced_mode”设置为“true”时才返回。
	Reviewer *string `json:"reviewer,omitempty"`

	// 开票人。 当“advanced_mode”设置为“true”时才返回。
	Issuer *string `json:"issuer,omitempty"`

	// 销售方发票专用章。 当“advanced_mode”设置为“true”时才返回。
	SellerSeal *[]string `json:"seller_seal,omitempty"`

	// 货物或应税劳务列表。
	ItemList *[]ItemList `json:"item_list,omitempty"`

	// 各个字段的置信度。 当“advanced_mode”设置为“true”时才返回。
	Confidence *interface{} `json:"confidence,omitempty"`

	// 文本框在原图位置。输出左上、右上、右下、左下四个点坐标。当“return_text_location”设置为“true”时才返回。
	TextLocation *interface{} `json:"text_location,omitempty"`
}

func (o VatInvoiceResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VatInvoiceResult struct{}"
	}

	return strings.Join([]string{"VatInvoiceResult", string(data)}, " ")
}
