package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MvsInvoiceResult struct {

	// 发票代码。
	Code *string `json:"code,omitempty"`

	// 发票号码。
	Number *string `json:"number,omitempty"`

	// 机打代码。
	MachinePrintedCode *string `json:"machine_printed_code,omitempty"`

	// 机打号码。
	MachinePrintedNumber *string `json:"machine_printed_number,omitempty"`

	// 开票日期。
	IssueDate *string `json:"issue_date,omitempty"`

	// 机器编号。
	MachineNumber *string `json:"machine_number,omitempty"`

	// 购买方名称。
	BuyerName *string `json:"buyer_name,omitempty"`

	// 购买方身份证号码/组织机构代码。
	BuyerOrganizationNumber *string `json:"buyer_organization_number,omitempty"`

	// 购买方纳税人识别号。
	BuyerId *string `json:"buyer_id,omitempty"`

	// 销货单位名称。
	SellerName *string `json:"seller_name,omitempty"`

	// 销售方电话。
	SellerPhone *string `json:"seller_phone,omitempty"`

	// 销售方纳税人识别号。
	SellerId *string `json:"seller_id,omitempty"`

	// 销售方账号。
	SellerAccount *string `json:"seller_account,omitempty"`

	// 销售方地址。
	SellerAddress *string `json:"seller_address,omitempty"`

	// 销售方开户行。
	SellerBank *string `json:"seller_bank,omitempty"`

	// 车辆类型。
	VehicleType *string `json:"vehicle_type,omitempty"`

	// 厂牌型号。
	BrandModel *string `json:"brand_model,omitempty"`

	// 产地。
	ManufacturingLocation *string `json:"manufacturing_location,omitempty"`

	// 合格证号。
	QualityCertificate *string `json:"quality_certificate,omitempty"`

	// 进口证明书号。
	ImportCertificate *string `json:"import_certificate,omitempty"`

	// 商检单号。
	InspectionNumber *string `json:"inspection_number,omitempty"`

	// 发动机号码。
	EngineNumber *string `json:"engine_number,omitempty"`

	// 车辆识别代号/车架号码。
	VehicleIdentificationNumber *string `json:"vehicle_identification_number,omitempty"`

	// 吨位。
	Tonnage *string `json:"tonnage,omitempty"`

	// 限乘人数。
	SeatingCapacity *string `json:"seating_capacity,omitempty"`

	// 主管税务机关。
	TaxAuthority *string `json:"tax_authority,omitempty"`

	// 主管税务机关代码。
	TaxAuthorityCode *string `json:"tax_authority_code,omitempty"`

	// 完税凭证号码。
	TaxPaymentReceipt *string `json:"tax_payment_receipt,omitempty"`

	// 增值税税率或征收率。
	TaxRate *string `json:"tax_rate,omitempty"`

	// 增值税税额。
	Tax *string `json:"tax,omitempty"`

	// 不含税价。
	TaxExclusivePrice *string `json:"tax_exclusive_price,omitempty"`

	// 价税合计。
	Total *string `json:"total,omitempty"`

	// 价税合计大写。
	TotalChinese *string `json:"total_chinese,omitempty"`

	// 税控码。
	FiscalCode *string `json:"fiscal_code,omitempty"`
}

func (o MvsInvoiceResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MvsInvoiceResult struct{}"
	}

	return strings.Join([]string{"MvsInvoiceResult", string(data)}, " ")
}
