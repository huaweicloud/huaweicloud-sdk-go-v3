package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

type MonthlyBillRecord struct {

	// 流水账单所在账期，东八区时间，格式为YYYY-MM。
	BillCycle *string `json:"bill_cycle,omitempty"`

	// 消费的客户账号ID。 如果是普通客户或者企业子客户查询消费记录，只能查询到客户自己的消费记录，且此处显示的是客户自己的客户ID。如果是企业主查询消费记录，可以查询到企业主以及企业子客户的消费记录，此处为消费的实际客户ID。如果是企业主自己的消费记录，则为企业主ID；如果是某个企业子客户的消费记录，则此处为企业子账号ID。
	CustomerId *string `json:"customer_id,omitempty"`

	// 云服务类型编码，例如OBS的云服务类型编码为“hws.service.type.obs”。您可以调用查询云服务类型列表接口获取。
	ServiceTypeCode *string `json:"service_type_code,omitempty"`

	// 资源类型编码，例如ECS的VM为“hws.resource.type.vm”。您可以调用查询资源类型列表接口获取。
	ResourceTypeCode *string `json:"resource_type_code,omitempty"`

	// 云服务类型名称。例如ECS的云服务类型名称为“弹性云服务器”。
	ServiceTypeName *string `json:"service_type_name,omitempty"`

	// 资源类型名称。例如ECS的资源类型名称为“云主机”。
	ResourceTypeName *string `json:"resource_type_name,omitempty"`

	// 云服务区编码，例如：“cn-north-1”。具体请参见地区和终端节点对应云服务的“区域”列的值。
	RegionCode *string `json:"region_code,omitempty"`

	// 企业项目标识（企业项目ID）。 default项目对应ID：0未归集（表示该云服务不支持企业项目管理能力）项目对应ID：null其余项目对应ID获取方法请参见[如何获取企业项目ID](https://support.huaweicloud.com/usermanual-em/zh-cn_topic_0126101490.html)。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 企业项目的名称。
	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`

	// 计费模式。 1：包年/包月3：按需10：预留实例11：节省计划
	ChargingMode *int32 `json:"charging_mode,omitempty"`

	// 消费时间。 计费模式为包年/包月和预留实例场景时为订单的支付时间。计费模式为按需场景时为话单的生/失效时间。
	ConsumeTime *string `json:"consume_time,omitempty"`

	// 交易时间，某条消费记录对应的扣费时间。
	TradeTime *string `json:"trade_time,omitempty"`

	// 服务商。 1：华为云2：云商店
	ProviderType *int32 `json:"provider_type,omitempty"`

	// 订单ID或交易ID，扣费维度的唯一标识。
	TradeId *string `json:"trade_id,omitempty"`

	// 唯一标识。 该字段为预留字段。
	Id *string `json:"id,omitempty"`

	// 账单类型。1：消费-新购 2：消费-续订 3：消费-变更 4：退款-退订 5：消费-使用 8：消费-自动续订 9：调账-补偿 14：消费-服务支持计划月末扣费 16：调账-扣费 18：消费-按月付费 20：退款-变更 23：消费-节省计划抵扣 24：退款-包年/包月转按需 103：消费-按年付费
	BillType *int32 `json:"bill_type,omitempty"`

	// 支付状态。 1：已支付2：未结清3：未结算
	Status *int32 `json:"status,omitempty"`

	// 官网价。单位：元。  说明： official_amount = official_discount_amount + erase_amount + consume_amount
	OfficialAmount *decimal.Decimal `json:"official_amount,omitempty"`

	// 折扣金额。单位：元。
	OfficialDiscountAmount *decimal.Decimal `json:"official_discount_amount,omitempty"`

	// 抹零金额。单位：元。
	EraseAmount *decimal.Decimal `json:"erase_amount,omitempty"`

	// 应付金额，包括现金券和储值卡和代金券金额。单位：元。  说明： consume_amount的值等于cash_amount，credit_amount，coupon_amount，flexipurchase_coupon_amount，stored_value_card_amount，bonus_amount，debt_amount，writeoff_amount的总和。
	ConsumeAmount *decimal.Decimal `json:"consume_amount,omitempty"`

	// 现金支付金额。单位：元.
	CashAmount *decimal.Decimal `json:"cash_amount,omitempty"`

	// 信用额度支付金额。单位：元。
	CreditAmount *decimal.Decimal `json:"credit_amount,omitempty"`

	// 代金券支付金额。单位：元。
	CouponAmount *decimal.Decimal `json:"coupon_amount,omitempty"`

	// 现金券支付金额。单位：元。
	FlexipurchaseCouponAmount *decimal.Decimal `json:"flexipurchase_coupon_amount,omitempty"`

	// 储值卡支付金额。单位：元。
	StoredValueCardAmount *decimal.Decimal `json:"stored_value_card_amount,omitempty"`

	// 奖励金支付金额（奖励金已经下线，目前用于现网客户未使用完的奖励金）。单位：元。
	BonusAmount *decimal.Decimal `json:"bonus_amount,omitempty"`

	// 欠费金额。单位：元。  说明： 对于月结客户，欠费金额即页面上的月度结算金额。
	DebtAmount *decimal.Decimal `json:"debt_amount,omitempty"`

	// 欠费核销金额。单位：元。
	WriteoffAmount *decimal.Decimal `json:"writeoff_amount,omitempty"`

	// 云服务区名称，例如：“华北-北京一”。具体请参见地区和终端节点对应云服务的“区域名称”列的值。
	RegionName *string `json:"region_name,omitempty"`
}

func (o MonthlyBillRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MonthlyBillRecord struct{}"
	}

	return strings.Join([]string{"MonthlyBillRecord", string(data)}, " ")
}
