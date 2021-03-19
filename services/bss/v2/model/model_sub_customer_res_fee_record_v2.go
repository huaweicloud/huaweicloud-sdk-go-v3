package model

import (
	"encoding/json"

	"strings"
)

type SubCustomerResFeeRecordV2 struct {
	// 消费记录对应的资源使用的开始时间。  说明： 按需有效，包年/包月该字段保留。

	EffectiveTime *string `json:"effective_time,omitempty"`
	// 消费记录对应的资源使用的结束时间。  说明： 按需有效，包年/包月该字段保留。

	ExpireTime *string `json:"expire_time,omitempty"`
	// 产品ID。

	ProductId *string `json:"product_id,omitempty"`
	// 产品名称。

	ProductName *string `json:"product_name,omitempty"`
	// 订单ID。  说明： 包年/包月资源的使用记录才有该字段，按需资源为空。

	OrderId *string `json:"order_id,omitempty"`
	// 消费金额，包括现金券和代金券金额，精确到小数点后2位。

	Amount *float64 `json:"amount,omitempty"`
	// 金额单位： 1：元

	MeasureId *int32 `json:"measure_id,omitempty"`
	// 资源的使用量类型，您可以调用查询使用量类型列表接口获取。

	UsageType *string `json:"usage_type,omitempty"`
	// 资源的使用量。

	Usage *float64 `json:"usage,omitempty"`
	// 资源的使用量单位，您可以调用查询使用量单位列表接口获取。

	UsageMeasureId *int32 `json:"usage_measure_id,omitempty"`
	// 套餐内使用量。

	FreeResourceUsage *float64 `json:"free_resource_usage,omitempty"`
	// 套餐内使用量单位，您可以调用查询使用量单位列表接口获取。

	FreeResourceMeasureId *int32 `json:"free_resource_measure_id,omitempty"`
	// 云服务类型编码，例如ECS的云服务类型编码为“hws.service.type.ec2”。您可以调用查询云服务类型列表接口获取。

	CloudServiceType *string `json:"cloud_service_type,omitempty"`
	// 云服务区编码，例如：“cn-north-1”。具体请参见地区和终端节点对应云服务的“区域”列的值。

	Region *string `json:"region,omitempty"`
	// 资源类型编码，例如ECS的VM为“hws.resource.type.vm”。您可以调用查询资源类型列表接口获取。

	ResourceType *string `json:"resource_type,omitempty"`
	// 计费模式。 1 : 包年/包月3：按需10：预留实例

	ChargeMode *string `json:"charge_mode,omitempty"`
	// 资源标签。

	ResourceTag *string `json:"resource_tag,omitempty"`
	// 资源名称。

	ResourceName *string `json:"resource_name,omitempty"`
	// 资源ID。

	ResourceId *string `json:"resource_id,omitempty"`
	// 账单类型。 1：消费-新购2：消费-续订3：消费-变更4：退款-退订5：消费-使用8：消费-自动续订9：调账-补偿14：消费-服务支持计划月末扣费16：调账-扣费

	BillType *int32 `json:"bill_type,omitempty"`
	// 周期类型： 19：年20：月24：天25：小时5：一次性  说明： charge_mode取值为“1：包年/包月资源”或“10：预留实例”时该参数才有值；charge_mode取值为“3：按需资源”时该参数为空。

	PeriodType *string `json:"period_type,omitempty"`
	// 产品的规格描述。例如：“普通IO|100.0GB”。

	ProductSpecDesc *string `json:"product_spec_desc,omitempty"`
	// 预留实例使用量。

	RiUsage *float64 `json:"ri_usage,omitempty"`
	// 预留实例使用量单位。

	RiUsageMeasureId *int32 `json:"ri_usage_measure_id,omitempty"`
	// 官网价。

	OfficialAmount *float64 `json:"official_amount,omitempty"`
	// 折扣金额

	DiscountAmount *float64 `json:"discount_amount,omitempty"`
	// 现金支付金额

	CashAmount *float64 `json:"cash_amount,omitempty"`
	// 信用额度支付金额。

	CreditAmount *float64 `json:"credit_amount,omitempty"`
	// 代金券支付金额。

	CouponAmount *float64 `json:"coupon_amount,omitempty"`
	// 现金券支付金额。

	FlexipurchaseCouponAmount *float64 `json:"flexipurchase_coupon_amount,omitempty"`
	// 储值卡支付金额。

	StoredCardAmount *float64 `json:"stored_card_amount,omitempty"`
	// 奖励金支付金额（用于现网客户未使用完的奖励金）。

	BonusAmount *float64 `json:"bonus_amount,omitempty"`
	// 欠费金额。

	DebtAmount *float64 `json:"debt_amount,omitempty"`
	// 欠费核销金额。

	AdjustmentAmount *float64 `json:"adjustment_amount,omitempty"`
	// 产品的实例大小，仅线性产品有效。 线性产品为包括硬盘，带宽等在订购时需要指定大小的产品。例如硬盘在订购时需选择10G、20G等不同大小。

	SpecSize *float64 `json:"spec_size,omitempty"`
	// 产品实例大小的单位，仅线性产品有该字段。您可以调用查询使用量单位列表接口获取。

	SpecSizeMeasureId *int32 `json:"spec_size_measure_id,omitempty"`
	// 云服务区名称，例如：“华北-北京一”。具体请参见地区和终端节点对应云服务的“区域名称”列的值。

	RegionName *string `json:"region_name,omitempty"`
}

func (o SubCustomerResFeeRecordV2) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SubCustomerResFeeRecordV2 struct{}"
	}

	return strings.Join([]string{"SubCustomerResFeeRecordV2", string(data)}, " ")
}
