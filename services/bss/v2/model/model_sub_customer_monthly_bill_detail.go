package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubCustomerMonthlyBillDetail struct {

	// 账期。 格式：YYYY-MM
	BillCycle *string `json:"bill_cycle,omitempty" xml:"bill_cycle"`

	// 客户账号ID。
	CustomerId *string `json:"customer_id,omitempty" xml:"customer_id"`

	// 子客户的关联类型： 1：顾问销售2：代售
	AssociationType *string `json:"association_type,omitempty" xml:"association_type"`

	// 云服务类型编码，例如OBS的云服务类型编码为“hws.service.type.obs”。您可以调用查询云服务类型列表接口获取。
	ServiceTypeCode *string `json:"service_type_code,omitempty" xml:"service_type_code"`

	// 资源类型编码，例如ECS的VM为“hws.resource.type.vm”。您可以调用查询资源类型列表接口获取。 ResourceType是CloudServiceType中的一种资源，CloudServiceType由多种ResourceType组合提供。
	ResourceTypeCode *string `json:"resource_type_code,omitempty" xml:"resource_type_code"`

	// 云服务类型名称。例如ECS的云服务类型名称为“弹性云服务器”。
	ServiceTypeName *string `json:"service_type_name,omitempty" xml:"service_type_name"`

	// 资源类型名称。例如ECS的资源类型名称为“云主机”。
	ResourceTypeName *string `json:"resource_type_name,omitempty" xml:"resource_type_name"`

	// 计费模式。 1：包周期3：按需10：预留实例
	ChargingMode *int32 `json:"charging_mode,omitempty" xml:"charging_mode"`

	// 交易时间，即某条消费记录对应的扣费时间。 示例：2020-11-17T06:43:38Z
	TradeTime *string `json:"trade_time,omitempty" xml:"trade_time"`

	// 订单ID或交易ID，扣费维度的唯一标识。 账单类型为1，2，3，4，8时为订单ID。其它场景下为交易ID。非月末扣费：应收ID月末扣费：账单ID
	TradeId *string `json:"trade_id,omitempty" xml:"trade_id"`

	// 账单类型。 1：消费-新购2：消费-续订3：消费-变更8：消费-自动续订5：消费-使用12：消费-按时计费4：退款-退订9：调账-补偿13：消费-退订手续费增收14：消费-服务支持计划月末补扣16：调账-扣费18：消费-按月付费20：退款-变更 100：退款-退订税金101：调账-补偿税金102：调账-扣费税金
	BillDetailType *int32 `json:"bill_detail_type,omitempty" xml:"bill_detail_type"`

	// 资源ID。
	ResourceId *string `json:"resource_id,omitempty" xml:"resource_id"`

	// 资源名称。
	ResourceName *string `json:"resource_name,omitempty" xml:"resource_name"`

	// 产品的规格描述。
	ProductSpecDesc *string `json:"product_spec_desc,omitempty" xml:"product_spec_desc"`

	// 云服务区编码，例如：“cn-north-1”。具体请参见地区和终端节点对应云服务的“区域”列的值。
	RegionCode *string `json:"region_code,omitempty" xml:"region_code"`

	// 产品ID。
	ProductId *string `json:"product_id,omitempty" xml:"product_id"`

	// 产品名称。
	ProductName *string `json:"product_name,omitempty" xml:"product_name"`

	// 资源标签。
	ResourceTag *string `json:"resource_tag,omitempty" xml:"resource_tag"`

	// 消费时间。 包周期和预留实例订购场景下为订单支付时间；按需场景下为话单生失效时间。 格式：YYYY-MM-DDThh:mm:ssZ
	ConsumeTime *string `json:"consume_time,omitempty" xml:"consume_time"`

	// 资源使用量的类型，您可以调用查询使用量类型列表接口获取。
	UsageType *string `json:"usage_type,omitempty" xml:"usage_type"`

	// 资源的使用量。
	UsageAmount *float64 `json:"usage_amount,omitempty" xml:"usage_amount"`

	// 资源使用量的度量单位，您可以调用查询度量单位列表接口获取。
	UsageMeasureId *int32 `json:"usage_measure_id,omitempty" xml:"usage_measure_id"`

	// 套餐内使用量。
	FreeResourceUsage *float64 `json:"free_resource_usage,omitempty" xml:"free_resource_usage"`

	// 套餐内使用量的度量单位，您可以调用查询度量单位列表接口获取。
	FreeResourceMeasureId *int32 `json:"free_resource_measure_id,omitempty" xml:"free_resource_measure_id"`

	// 预留实例使用量。
	RiUsage *float64 `json:"ri_usage,omitempty" xml:"ri_usage"`

	// 预留实例使用量单位。
	RiUsageMeasureId *int32 `json:"ri_usage_measure_id,omitempty" xml:"ri_usage_measure_id"`

	// 官网价。
	OfficialAmount *float64 `json:"official_amount,omitempty" xml:"official_amount"`

	// 对应官网价折扣金额。
	OfficialDiscountAmount *float64 `json:"official_discount_amount,omitempty" xml:"official_discount_amount"`

	// 应付金额。
	PaymentAmount *float64 `json:"payment_amount,omitempty" xml:"payment_amount"`

	// 现金支付金额。
	CashAmount *float64 `json:"cash_amount,omitempty" xml:"cash_amount"`

	// 信用额度支付金额。
	CreditAmount *float64 `json:"credit_amount,omitempty" xml:"credit_amount"`

	// 代金券支付金额。
	CouponAmount *float64 `json:"coupon_amount,omitempty" xml:"coupon_amount"`

	// 现金券支付金额。
	FlexipurchaseCouponAmount *float64 `json:"flexipurchase_coupon_amount,omitempty" xml:"flexipurchase_coupon_amount"`

	// 储值卡支付金额。
	StoredValueCardAmount *float64 `json:"stored_value_card_amount,omitempty" xml:"stored_value_card_amount"`

	// 欠费金额。
	DebtAmount *float64 `json:"debt_amount,omitempty" xml:"debt_amount"`

	// 欠费核销金额。
	WriteoffAmount *float64 `json:"writeoff_amount,omitempty" xml:"writeoff_amount"`

	// 周期类型： 19：年20：月24：天25：小时5：一次性
	PeriodType *int32 `json:"period_type,omitempty" xml:"period_type"`

	// 客户经理标识。
	AccountManagerId *string `json:"account_manager_id,omitempty" xml:"account_manager_id"`

	// 关联的经销商ID。
	PartnerId *string `json:"partner_id,omitempty" xml:"partner_id"`

	// 云服务区名称，例如：“华北-北京一”。具体请参见地区和终端节点对应云服务的“区域名称”列的值。
	RegionName *string `json:"region_name,omitempty" xml:"region_name"`
}

func (o SubCustomerMonthlyBillDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubCustomerMonthlyBillDetail struct{}"
	}

	return strings.Join([]string{"SubCustomerMonthlyBillDetail", string(data)}, " ")
}
