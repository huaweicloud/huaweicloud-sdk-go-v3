package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MvsInvoiceResult struct {

	// 发票代码。
	Code *string `json:"code,omitempty" xml:"code"`

	// 发票号码。
	Number *string `json:"number,omitempty" xml:"number"`

	// 机打代码。
	MachinePrintedCode *string `json:"machine_printed_code,omitempty" xml:"machine_printed_code"`

	// 机打号码。
	MachinePrintedNumber *string `json:"machine_printed_number,omitempty" xml:"machine_printed_number"`

	// 开票日期。
	IssueDate *string `json:"issue_date,omitempty" xml:"issue_date"`

	// 机器编号。
	MachineNumber *string `json:"machine_number,omitempty" xml:"machine_number"`

	// 购买方名称。
	BuyerName *string `json:"buyer_name,omitempty" xml:"buyer_name"`

	// 购买方身份证号码/组织机构代码。
	BuyerOrganizationNumber *string `json:"buyer_organization_number,omitempty" xml:"buyer_organization_number"`

	// 购买方纳税人识别号。
	BuyerId *string `json:"buyer_id,omitempty" xml:"buyer_id"`

	// 销货单位名称。
	SellerName *string `json:"seller_name,omitempty" xml:"seller_name"`

	// 销售方电话。
	SellerPhone *string `json:"seller_phone,omitempty" xml:"seller_phone"`

	// 销售方纳税人识别号。
	SellerId *string `json:"seller_id,omitempty" xml:"seller_id"`

	// 销售方账号。
	SellerAccount *string `json:"seller_account,omitempty" xml:"seller_account"`

	// 销售方地址。
	SellerAddress *string `json:"seller_address,omitempty" xml:"seller_address"`

	// 销售方开户行。
	SellerBank *string `json:"seller_bank,omitempty" xml:"seller_bank"`

	// 车辆类型。
	VehicleType *string `json:"vehicle_type,omitempty" xml:"vehicle_type"`

	// 厂牌型号。
	BrandModel *string `json:"brand_model,omitempty" xml:"brand_model"`

	// 产地。
	ManufacturingLocation *string `json:"manufacturing_location,omitempty" xml:"manufacturing_location"`

	// 合格证号。
	QualityCertificate *string `json:"quality_certificate,omitempty" xml:"quality_certificate"`

	// 进口证明书号。
	ImportCertificate *string `json:"import_certificate,omitempty" xml:"import_certificate"`

	// 商检单号。
	InspectionNumber *string `json:"inspection_number,omitempty" xml:"inspection_number"`

	// 发动机号码。
	EngineNumber *string `json:"engine_number,omitempty" xml:"engine_number"`

	// 车辆识别代号/车架号码。
	VehicleIdentificationNumber *string `json:"vehicle_identification_number,omitempty" xml:"vehicle_identification_number"`

	// 吨位。
	Tonnage *string `json:"tonnage,omitempty" xml:"tonnage"`

	// 限乘人数。
	SeatingCapacity *string `json:"seating_capacity,omitempty" xml:"seating_capacity"`

	// 主管税务机关。
	TaxAuthority *string `json:"tax_authority,omitempty" xml:"tax_authority"`

	// 主管税务机关代码。
	TaxAuthorityCode *string `json:"tax_authority_code,omitempty" xml:"tax_authority_code"`

	// 完税凭证号码。
	TaxPaymentReceipt *string `json:"tax_payment_receipt,omitempty" xml:"tax_payment_receipt"`

	// 增值税税率或征收率。
	TaxRate *string `json:"tax_rate,omitempty" xml:"tax_rate"`

	// 增值税税额。
	Tax *string `json:"tax,omitempty" xml:"tax"`

	// 不含税价。
	TaxExclusivePrice *string `json:"tax_exclusive_price,omitempty" xml:"tax_exclusive_price"`

	// 价税合计。
	Total *string `json:"total,omitempty" xml:"total"`

	// 价税合计大写。
	TotalChinese *string `json:"total_chinese,omitempty" xml:"total_chinese"`

	// 税控码。
	FiscalCode *string `json:"fiscal_code,omitempty" xml:"fiscal_code"`
}

func (o MvsInvoiceResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MvsInvoiceResult struct{}"
	}

	return strings.Join([]string{"MvsInvoiceResult", string(data)}, " ")
}
