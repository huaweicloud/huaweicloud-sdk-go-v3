package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

type ResFeeRecordV2 struct {

	// 资源消费记录的日期。 格式：YYYY-MM-DD。按照东八区截取。
	BillDate *string `json:"bill_date,omitempty"`

	// 账单类型。 1：消费-新购2：消费-续订3：消费-变更4：退款-退订5：消费-使用8：消费-自动续订9：调账-补偿14：消费-服务支持计划月末扣费15：消费-税金16：调账-扣费17：消费-保底差额 说明： 保底差额=客户签约保底合同后，如果没有达到保底消费，客户需要补交的费用，仅限于直销或者伙伴顾问销售类子客户，且为后付费用户。 20：退款-变更23：消费-节省计划抵扣24：退款-包年/包月转按需 100：退款-退订税金101：调账-补偿税金102：调账-扣费税金|
	BillType *int32 `json:"bill_type,omitempty"`

	// 消费的客户账号ID。 如果是普通客户或者企业子查询消费记录，只能查询到自身的消费记录，则这个地方显示的是自身的客户ID如果是企业主查询消费记录，可以查询到自身以及企业子的消费记录，这个地方是消费的实际客户ID，如果是企业主自身消费，为企业主ID，如果这条消费记录是某个企业子客户的消费，这个地方的ID是企业子账号ID。
	CustomerId *string `json:"customer_id,omitempty"`

	// 云服务区编码，例如：“ap-southeast-1”。具体请参见地区和终端节点对应云服务的“区域”列的值。
	Region *string `json:"region,omitempty"`

	// 云服务区名称，例如：“中国-香港”。具体请参见地区和终端节点对应云服务的“区域名称”列的值。
	RegionName *string `json:"region_name,omitempty"`

	// 云服务类型编码，例如OBS的云服务类型编码为“hws.service.type.obs”。您可以调用查询云服务类型列表接口获取。
	CloudServiceType *string `json:"cloud_service_type,omitempty"`

	// 资源类型编码，例如ECS的VM为“hws.resource.type.vm”。
	ResourceType *string `json:"resource_type,omitempty"`

	// 云服务类型名称。例如ECS的云服务类型名称为“弹性云服务器”。
	CloudServiceTypeName *string `json:"cloud_service_type_name,omitempty"`

	// 资源类型名称。例如ECS的资源类型名称为“云主机”。
	ResourceTypeName *string `json:"resource_type_name,omitempty"`

	// 费用对应的资源使用的开始时间，按需有效，包年/包月该字段保留。
	EffectiveTime *string `json:"effective_time,omitempty"`

	// 费用对应的资源使用的结束时间，按需有效，包年/包月该字段保留。
	ExpireTime *string `json:"expire_time,omitempty"`

	// 资源ID。
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源名称。
	ResourceName *string `json:"resource_name,omitempty"`

	// 资源标签。
	ResourceTag *string `json:"resource_tag,omitempty"`

	// 产品ID。
	ProductId *string `json:"product_id,omitempty"`

	// 产品名称。
	ProductName *string `json:"product_name,omitempty"`

	// 产品的规格描述。
	ProductSpecDesc *string `json:"product_spec_desc,omitempty"`

	// SKU编码，在账单中唯一标识一个资源的规格。
	SkuCode *string `json:"sku_code,omitempty"`

	// 产品的实例大小，仅线性产品有效。  说明： 线性产品是指订购时需要指定大小的产品。例如硬盘在订购时需选择10G、20G等不同大小规格。
	SpecSize *decimal.Decimal `json:"spec_size,omitempty"`

	// 产品实例大小的单位，仅线性产品有该字段。 您可以调用查询度量单位列表接口获取。
	SpecSizeMeasureId *int32 `json:"spec_size_measure_id,omitempty"`

	// 订单ID或交易ID，扣费维度的唯一标识。
	TradeId *string `json:"trade_id,omitempty"`

	// |参数名称：唯一标识。|
	Id *string `json:"id,omitempty"`

	// 交易时间。
	TradeTime *string `json:"trade_time,omitempty"`

	// 企业项目标识（企业项目ID）。 default项目对应ID：0未归集（表示该云服务不支持企业项目管理能力）项目对应ID：null
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 企业项目的名称。
	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`

	// 计费模式。 1：包年/包月3：按需10：预留实例
	ChargeMode *string `json:"charge_mode,omitempty"`

	// 订单ID。  说明： 包年/包月资源的使用记录才有该字段，按需资源则为空。
	OrderId *string `json:"order_id,omitempty"`

	// 周期类型： 19：年20：月24：天25：小时5：一次性
	PeriodType *string `json:"period_type,omitempty"`

	// 资源使用量的类型，您可以调用查询使用量类型列表接口获取。
	UsageType *string `json:"usage_type,omitempty"`

	// 资源的使用量。
	Usage *decimal.Decimal `json:"usage,omitempty"`

	// 资源使用量的度量单位，您可以调用查询度量单位列表接口获取。
	UsageMeasureId *int32 `json:"usage_measure_id,omitempty"`

	// 套餐内使用量。
	FreeResourceUsage *decimal.Decimal `json:"free_resource_usage,omitempty"`

	// 套餐内使用量的度量单位，您可以调用查询度量单位列表接口获取。
	FreeResourceMeasureId *int32 `json:"free_resource_measure_id,omitempty"`

	// 预留实例使用量。
	RiUsage *decimal.Decimal `json:"ri_usage,omitempty"`

	// 预留实例使用量单位。
	RiUsageMeasureId *int32 `json:"ri_usage_measure_id,omitempty"`

	// 产品的单价。 按需产品的单价，只有简单定价，不分档的场景会返回。 包周期产品的单价，只有包周期的如下场景会返回：包周期订购/续订/降配/升配/扩容简单定价，不分档 预留实例的单价，只有如下场景下会返回：订购/续订/降配/升配/扩容/按时计费简单定价，不分档
	UnitPrice *decimal.Decimal `json:"unit_price,omitempty"`

	// 产品的单价单位。 线性产品的单价单位为“元/{线性单位}/月”或“元/{线性单位}/小时”等。非线性产品的单价单位为“元/月”或“元/小时”等。  说明： “线性单位”为线性产品（即订购时需要指定大小的产品）的大小的单位，比如硬盘的线性单位为GB，带宽的线性单位为Mbps。
	Unit *string `json:"unit,omitempty"`

	// 官网价，华为云商品在官网上未叠加应用商务折扣、促销折扣等优惠的销售价格。
	OfficialAmount *decimal.Decimal `json:"official_amount,omitempty"`

	// 优惠金额，用户使用云服务享受折扣优惠如商务折扣、伙伴授予折扣以及促销优惠等减免的金额。
	DiscountAmount *decimal.Decimal `json:"discount_amount,omitempty"`

	// 应付金额，用户使用云服务享受折扣优惠后需要支付的费用金额，包括代金券金额，精确到小数点后8位。  说明： amount的值等于cash_amount，credit_amount，coupon_amount，flexipurchase_coupon_amount，stored_card_amount，bonus_amount，debt_amount，adjustment_amount的总和。
	Amount *decimal.Decimal `json:"amount,omitempty"`

	// 现金支付金额。
	CashAmount *decimal.Decimal `json:"cash_amount,omitempty"`

	// 信用额度支付金额。
	CreditAmount *decimal.Decimal `json:"credit_amount,omitempty"`

	// 代金券支付金额。
	CouponAmount *decimal.Decimal `json:"coupon_amount,omitempty"`

	// 现金券支付金额。
	FlexipurchaseCouponAmount *decimal.Decimal `json:"flexipurchase_coupon_amount,omitempty"`

	// 储值卡支付金额。
	StoredCardAmount *decimal.Decimal `json:"stored_card_amount,omitempty"`

	// 奖励金支付金额（用于现网客户未使用完的奖励金）。
	BonusAmount *decimal.Decimal `json:"bonus_amount,omitempty"`

	// 欠费金额。
	DebtAmount *decimal.Decimal `json:"debt_amount,omitempty"`

	// 欠费核销金额。
	AdjustmentAmount *decimal.Decimal `json:"adjustment_amount,omitempty"`

	// 金额单位。 1：元
	MeasureId *int32 `json:"measure_id,omitempty"`

	// 实付金额计算公式。当前只包含如下场景： 按需简单定价 按需线性定价 包年包月新购和续费的简单定价 包年包月新购和续费的线性定价  说明： 实付金额计算公式得到的金额值等于amount - coupon_amount的差值。
	Formula *string `json:"formula,omitempty"`

	// 该字段为预留字段。
	SubServiceTypeCode *string `json:"sub_service_type_code,omitempty"`

	// 该字段为预留字段。
	SubServiceTypeName *string `json:"sub_service_type_name,omitempty"`

	// 该字段为预留字段。
	SubResourceTypeCode *string `json:"sub_resource_type_code,omitempty"`

	// 该字段为预留字段。
	SubResourceTypeName *string `json:"sub_resource_type_name,omitempty"`

	// 该字段为预留字段。
	SubResourceId *string `json:"sub_resource_id,omitempty"`

	// 该字段为预留字段。
	SubResourceName *string `json:"sub_resource_name,omitempty"`
}

func (o ResFeeRecordV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResFeeRecordV2 struct{}"
	}

	return strings.Join([]string{"ResFeeRecordV2", string(data)}, " ")
}
