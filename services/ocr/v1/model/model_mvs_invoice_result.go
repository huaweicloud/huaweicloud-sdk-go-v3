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

	// 购买方名称、二手车买方单位/个人
	BuyerName *string `json:"buyer_name,omitempty"`

	// 购买方身份证号码/组织机构代码。
	BuyerOrganizationNumber *string `json:"buyer_organization_number,omitempty"`

	// 购买方纳税人识别号、二手车买方单位代码/身份证号
	BuyerId *string `json:"buyer_id,omitempty"`

	// 二手车买方单位/个人住址。 当请求参数\"type\"设置为\"auto\"或\"used\"时才返回。
	BuyerAddress *string `json:"buyer_address,omitempty"`

	// 二手车买方单位/个人电话。 当请求参数\"type\"设置为\"auto\"或\"used\"时才返回。
	BuyerPhone *string `json:"buyer_phone,omitempty"`

	// 销货单位名称、二手车卖方单位/个人
	SellerName *string `json:"seller_name,omitempty"`

	// 销售方电话、二手车卖方电话
	SellerPhone *string `json:"seller_phone,omitempty"`

	// 销售方纳税人识别号、二手车卖方单位代码/身份证号
	SellerId *string `json:"seller_id,omitempty"`

	// 销售方账号。
	SellerAccount *string `json:"seller_account,omitempty"`

	// 销售方地址、二手车卖方单位/个人地址
	SellerAddress *string `json:"seller_address,omitempty"`

	// 二手车车牌照号。 当请求参数\"type\"设置为\"auto\"或\"used\"时才返回。
	LicencePlateNumber *string `json:"licence_plate_number,omitempty"`

	// 二手车登记证号。 当请求参数\"type\"设置为\"auto\"或\"used\"时才返回。
	RegistrationNumber *string `json:"registration_number,omitempty"`

	// 二手车转入地车管所名称。 当请求参数\"type\"设置为\"auto\"或\"used\"时才返回。
	DeptMotorVehicles *string `json:"dept_motor_vehicles,omitempty"`

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

	// 价税合计、二手车车价合计（小写）
	Total *string `json:"total,omitempty"`

	// 价税合计大写、二手车车价合计（大写）
	TotalChinese *string `json:"total_chinese,omitempty"`

	// 税控码。
	FiscalCode *string `json:"fiscal_code,omitempty"`

	// 二手车经营拍卖单位名称。 当请求参数\"type\"设置为\"auto\"或\"used\"时才返回。
	AuctionOrgName *string `json:"auction_org_name,omitempty"`

	// 二手车经营拍卖单位地址。 当请求参数\"type\"设置为\"auto\"或\"used\"时才返回。
	AuctionOrgAddress *string `json:"auction_org_address,omitempty"`

	// 二手车经营拍卖单位纳税人识别号。 当请求参数\"type\"设置为\"auto\"或\"used\"时才返回。
	AuctionOrgId *string `json:"auction_org_id,omitempty"`

	// 二手车经营拍卖单位银行和账号。 当请求参数\"type\"设置为\"auto\"或\"used\"时才返回。
	AuctionOrgBankAccount *string `json:"auction_org_bank_account,omitempty"`

	// 二手车经营拍卖单位电话。 当请求参数\"type\"设置为\"auto\"或\"used\"时才返回。
	AuctionOrgPhone *string `json:"auction_org_phone,omitempty"`

	// 二手车市场名称。 当请求参数\"type\"设置为\"auto\"或\"used\"时才返回。
	UsedVehicleMarketName *string `json:"used_vehicle_market_name,omitempty"`

	// 二手车市场纳税人识别号。 当请求参数\"type\"设置为\"auto\"或\"used\"时才返回。
	UsedVehicleMarketId *string `json:"used_vehicle_market_id,omitempty"`

	// 二手车市场地址。 当请求参数\"type\"设置为\"auto\"或\"used\"时才返回。
	UsedVehicleMarketAddress *string `json:"used_vehicle_market_address,omitempty"`

	// 二手车市场银行和账号。 当请求参数\"type\"设置为\"auto\"或\"used\"时才返回。
	UsedVehicleMarketBankAccount *string `json:"used_vehicle_market_bank_account,omitempty"`

	// 二手车市场电话。 当请求参数\"type\"设置为\"auto\"或\"used\"时才返回。
	UsedVehicleMarketPhone *string `json:"used_vehicle_market_phone,omitempty"`

	// 二手车反向开具标识，True表示是反向开具发票，False表示不是反向开具发票。 当请求参数\"type\"设置为\"auto\"或\"used\"时才返回。
	ReverseIssue *bool `json:"reverse_issue,omitempty"`

	// 备注
	Remark *string `json:"remark,omitempty"`

	// 开票人
	DrawerName *string `json:"drawer_name,omitempty"`

	// 枚举值，机动车销售统一发票或者二手车销售统一发票。 当入参中包含type时返回。
	Type *string `json:"type,omitempty"`

	// 检测框对象，内部为字段-框坐标对。如 code:[[x0, y0],[x1,y1],[x2,y2],[x3,y3]], 点的顺序是左上角、右上角、右下角、左下角。如果原图找不到字段，返回空列表。
	TextLocation *interface{} `json:"text_location,omitempty"`

	// 字段文字内容置信度，内容为字段-字符置信度对,如code:0.9999。这个数值为字段中每个字符置信度，格式为fp32，保留四位。若字段不存在则返回0.0。
	Confidence *interface{} `json:"confidence,omitempty"`
}

func (o MvsInvoiceResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MvsInvoiceResult struct{}"
	}

	return strings.Join([]string{"MvsInvoiceResult", string(data)}, " ")
}
