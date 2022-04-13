package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NvlCostAnalysedBillDetail struct {
	// 查询分摊成本的月份。 格式为YYYY-MM，按照东八区截取。

	SharedMonth *string `json:"shared_month,omitempty"`
	// 账期。 格式：YYYY-MM。按照东八区截取。

	BillCycle *string `json:"bill_cycle,omitempty"`
	// 账单类型。 1：消费-新购2：消费-续订3：消费-变更4：退款-退订5：消费-使用8：消费-自动续订9：调账-补偿14：消费-服务支持计划月末扣费16：调账-扣费18：消费-按月付费20：退款-变更

	BillType *int32 `json:"bill_type,omitempty"`
	// 消费的客户账号ID。 如果是普通客户或者企业子查询消费记录，只能查询到自身的消费记录，则这个地方显示的是自身的客户ID。如果是企业主查询消费记录，可以查询到自身以及企业子的消费记录，这个地方是消费的实际客户ID，如果是企业主自身消费，为企业主ID，如果这条消费记录是某个企业子客户的消费，这个地方的ID是企业子账号ID。

	CustomerId *string `json:"customer_id,omitempty"`
	// 云服务区编码，例如：“cn-north-1”。具体请参见地区和终端节点对应云服务的“区域”列的值。

	RegionCode *string `json:"region_code,omitempty"`
	// 云服务区名称，例如：“华北-北京一”。具体请参见地区和终端节点对应云服务的“区域名称”列的值。

	RegionName *string `json:"region_name,omitempty"`
	// 云服务类型编码，例如OBS的云服务类型编码为“hws.service.type.obs”。您可以调用查询云服务类型列表接口获取。

	ServiceTypeCode *string `json:"service_type_code,omitempty"`
	// 资源类型编码，例如ECS的VM为“hws.resource.type.vm”。您可以调用查询资源类型列表接口获取。

	ResourceTypeCode *string `json:"resource_type_code,omitempty"`
	// 云服务类型名称。例如ECS的云服务类型名称为“弹性云服务器”。

	ServiceTypeName *string `json:"service_type_name,omitempty"`
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
	// 产品的规格描述。

	ProductSpecDesc *string `json:"product_spec_desc,omitempty"`
	// 企业项目标识（企业项目ID）。 default项目对应ID：0未归集（表示该云服务不支持企业项目管理能力）项目对应ID：-1其余项目对应ID获取方法请参见如何获取企业项目ID。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 企业项目的名称。

	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`
	// 计费模式。 1：包年/包月3：按需10：预留实例

	ChargingMode *int32 `json:"charging_mode,omitempty"`
	// 订单ID。  说明： 包年/包月资源的使用记录才有该字段，按需资源则为空。

	OrderId *string `json:"order_id,omitempty"`
	// 周期类型。 19：年20：月24：天25：小时5：一次性

	PeriodType *int32 `json:"period_type,omitempty"`
	// 资源使用量的类型，您可以调用查询使用量类型列表接口获取。

	UsageType *string `json:"usage_type,omitempty"`
	// 资源的使用量。

	Usage *float64 `json:"usage,omitempty"`
	// 资源使用量的度量单位，您可以调用查询度量单位列表接口获取。

	UsageMeasureId *int32 `json:"usage_measure_id,omitempty"`
	// 套餐内使用量。

	FreeResourceUsage *float64 `json:"free_resource_usage,omitempty"`
	// 套餐内使用量的度量单位，您可以调用查询度量单位列表接口获取。

	FreeResourceMeasureId *int32 `json:"free_resource_measure_id,omitempty"`
	// 预留实例使用量。

	RiUsage *float64 `json:"ri_usage,omitempty"`
	// 预留实例使用量单位。

	RiUsageMeasureId *int32 `json:"ri_usage_measure_id,omitempty"`
	// 消费金额（应付金额）。 消费金额=期初已分摊金额+当月分摊金额+期末未分摊金额

	ConsumeAmount *float64 `json:"consume_amount,omitempty"`
	// 期初已分摊金额。  说明： 包周期和预留实例预付时有效；计费类型为按需，预留实例为按时计费时该值为0。

	PastMonthsAmortizedAmount *float64 `json:"past_months_amortized_amount,omitempty"`
	// 当月分摊金额。 当月分摊金额=现金分摊金额+信用额度分摊金额+代金券分摊金额+现金券分摊金额+储值卡分摊金额+奖励金分摊金额

	CurrentMonthAmortizedAmount *float64 `json:"current_month_amortized_amount,omitempty"`
	// 期末未分摊金额。月度成本分摊时，当月以后还未分摊的金额。  说明： 包周期和预留实例预付时有效；计费类型为按需，预留实例为按时计费时该值为0。

	FutureMonthsAmortizedAmount *float64 `json:"future_months_amortized_amount,omitempty"`
	// 月度成本分摊时，当月已分摊金额中包含的现金金额。

	AmortizedCashAmount *float64 `json:"amortized_cash_amount,omitempty"`
	// 月度成本分摊时，当月已分摊金额中包含的信用额度分摊金额。

	AmortizedCreditAmount *float64 `json:"amortized_credit_amount,omitempty"`
	// 月度成本分摊时，当月已分摊金额中包含的代金券分摊金额。

	AmortizedCouponAmount *float64 `json:"amortized_coupon_amount,omitempty"`
	// 月度成本分摊时，当月已分摊金额中包含的现金券分摊金额。

	AmortizedFlexipurchaseCouponAmount *float64 `json:"amortized_flexipurchase_coupon_amount,omitempty"`
	// 月度成本分摊时，当月已分摊金额中包含的储值卡分摊金额。

	AmortizedStoredValueCardAmount *float64 `json:"amortized_stored_value_card_amount,omitempty"`
	// 月度成本分摊时，当月已分摊金额中包含的奖励金分摊金额（用于现网未清干净的奖励金）。

	AmortizedBonusAmount *float64 `json:"amortized_bonus_amount,omitempty"`
}

func (o NvlCostAnalysedBillDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NvlCostAnalysedBillDetail struct{}"
	}

	return strings.Join([]string{"NvlCostAnalysedBillDetail", string(data)}, " ")
}
