package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NvlCostAnalysedBillDetail struct {
	// |参数名称：分摊月，格式为YYYY-MM，按照东八区截取。| |参数约束及描述：|

	SharedMonth *string `json:"shared_month,omitempty"`
	// |参数名称：账期。格式：YYYY-MM。按照东八区截取。| |参数约束及描述：|

	BillCycle *string `json:"bill_cycle,omitempty"`
	// |参数名称：账单类型。1：消费-新购2：消费-续订3：消费-变更4：退款-退订5：消费-使用8：消费-自动续订9：调账-补偿12：消费-按时计费13：消费-退订手续费14：消费-服务支持计划月末扣费15：消费-税金16：调账-扣费17：消费-保底差额保底差额=客户签约保底合同后，如果没有达到保底消费，客户需要补交的费用，仅限于直销或者伙伴推荐类子客户，且为后付费用户。100：退款-退订税金101：调账-补偿税金102：调账-扣费税金| |参数的约束及描述：|

	BillType *int32 `json:"bill_type,omitempty"`
	// |参数名称：消费的客户账号ID。如果是普通客户或者企业子查询消费记录，只能查询到自身的消费记录，则这个地方显示的是自身的客户ID如果是企业主查询消费记录，可以查询到自身以及企业子的消费记录，这个地方是消费的实际客户ID，如果是企业主自身消费，为企业主ID，如果这条消费记录是某个企业子客户的消费，这个地方的ID是企业子账号ID。| |参数约束及描述：|

	CustomerId *string `json:"customer_id,omitempty"`
	// |参数名称：云服务区编码，例如：“cn-north-1”。具体请参见地区和终端节点对应云服务的“区域”列的值。| |参数约束及描述：|

	RegionCode *string `json:"region_code,omitempty"`
	// |参数名称：云服务区名称，例如：“华北-北京一”。具体请参见地区和终端节点对应云服务的“区域名称”列的值。| |参数约束及描述：|

	RegionName *string `json:"region_name,omitempty"`
	// |参数名称：云服务类型编码，例如ECS的云服务类型编码为“hws.service.type.ec2”。您可以调用查询云服务类型列表接口获取。| |参数约束及描述：|

	ServiceTypeCode *string `json:"service_type_code,omitempty"`
	// |参数名称：资源类型编码，例如ECS的VM为“hws.resource.type.vm”。您可以调用查询资源类型列表接口获取。| |参数约束及描述：|

	ResourceTypeCode *string `json:"resource_type_code,omitempty"`
	// |参数名称：费用对应的资源使用的开始时间，按需为资源使用时间，包周期为下单后，对应那一单的资源生效时间。| |参数约束及描述：|

	EffectiveTime *string `json:"effective_time,omitempty"`
	// |参数名称：费用对应的资源使用的结束时间，按需为资源使用时间，包周期为下单后，对应那一单的资源失效时间。| |参数约束及描述：|

	ExpireTime *string `json:"expire_time,omitempty"`
	// |参数名称：资源ID。| |参数约束及描述：|

	ResourceId *string `json:"resource_id,omitempty"`
	// |参数名称：资源名称。| |参数约束及描述：|

	ResourceName *string `json:"resource_name,omitempty"`
	// |参数名称：资源标签。| |参数约束及描述：|

	ResourceTag *string `json:"resource_tag,omitempty"`
	// |参数名称：产品的规格描述。| |参数约束及描述：|

	ProductSpecDesc *string `json:"product_spec_desc,omitempty"`
	// |参数名称：企业项目标识（企业项目ID）。default项目对应ID：0未归集（表示该云服务不支持企业项目管理能力）项目对应ID：-1其余项目对应ID获取方法请参见如何获取企业项目ID。| |参数约束及描述：|

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// |参数名称：企业项目的名称| |参数约束及描述：|

	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`
	// |参数名称：计费模式。1：包年/包月3：按需10：预留实例| |参数约束及描述：|

	ChargingMode *int32 `json:"charging_mode,omitempty"`
	// |参数名称：订单ID。包年/包月资源的使用记录才有该字段，按需资源则为空。| |参数约束及描述：|

	OrderId *string `json:"order_id,omitempty"`
	// |参数名称：周期类型：19：年20：月24：天25：小时5：一次性| |参数约束及描述：|

	PeriodType *int32 `json:"period_type,omitempty"`
	// |参数名称：资源使用量的类型，您可以调用查询使用量类型列表接口获取。| |参数约束及描述：|

	UsageType *string `json:"usage_type,omitempty"`
	// |参数名称：资源的使用量。| |参数的约束及描述：|

	Usage *float64 `json:"usage,omitempty"`
	// |参数名称：资源使用量的度量单位，您可以调用查询使用量单位列表接口获取。| |参数的约束及描述：|

	UsageMeasureId *int32 `json:"usage_measure_id,omitempty"`
	// |参数名称：套餐内使用量| |参数的约束及描述：

	FreeResourceUsage *float64 `json:"free_resource_usage,omitempty"`
	// |参数名称：套餐内使用量的度量单位，您可以调用查询使用量单位列表接口获取。| |参数的约束及描述：|

	FreeResourceMeasureId *int32 `json:"free_resource_measure_id,omitempty"`
	// |参数名称：预留实例使用量。| |参数的约束及描述：|

	RiUsage *float64 `json:"ri_usage,omitempty"`
	// |参数名称：预留实例使用量单位。| |参数的约束及描述：|

	RiUsageMeasureId *int32 `json:"ri_usage_measure_id,omitempty"`
	// |参数名称：消费金额（应付金额）| |参数的约束及描述：|

	ConsumeAmount *float64 `json:"consume_amount,omitempty"`
	// |参数名称：期初已分摊金额（包周期和预留实例预付时有效，计费类型为按需，预留实例按时计费时为0）| |参数的约束及描述：|

	PastMonthsAmortizedAmount *float64 `json:"past_months_amortized_amount,omitempty"`
	// |参数名称：当月分摊金额| |参数的约束及描述：|

	CurrentMonthAmortizedAmount *float64 `json:"current_month_amortized_amount,omitempty"`
	// |参数名称：期末未分摊金额（包周期和预留实例预付时有效，计费类型为按需，预留实例按时计费时为0）| |参数的约束及描述：|

	FutureMonthsAmortizedAmount *float64 `json:"future_months_amortized_amount,omitempty"`
	// |参数名称：现金分摊金额| |参数的约束及描述：|

	AmortizedCashAmount *float64 `json:"amortized_cash_amount,omitempty"`
	// |参数名称：信用额度分摊金额| |参数的约束及描述：|

	AmortizedCreditAmount *float64 `json:"amortized_credit_amount,omitempty"`
	// |参数名称：代金券分摊金额| |参数的约束及描述：|

	AmortizedCouponAmount *float64 `json:"amortized_coupon_amount,omitempty"`
	// |参数名称：现金券分摊金额| |参数的约束及描述：|

	AmortizedFlexipurchaseCouponAmount *float64 `json:"amortized_flexipurchase_coupon_amount,omitempty"`
	// |参数名称：储值卡分摊金额| |参数的约束及描述：|

	AmortizedStoredValueCardAmount *float64 `json:"amortized_stored_value_card_amount,omitempty"`
	// |参数名称：奖励金分摊金额（用于现网未清干净的奖励金）| |参数的约束及描述：|

	AmortizedBonusAmount *float64 `json:"amortized_bonus_amount,omitempty"`
}

func (o NvlCostAnalysedBillDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NvlCostAnalysedBillDetail struct{}"
	}

	return strings.Join([]string{"NvlCostAnalysedBillDetail", string(data)}, " ")
}
