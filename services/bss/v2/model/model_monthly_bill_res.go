package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MonthlyBillRes struct {
	// 资源详单数据所在账期，格式为YYYY-MM。 例如2020-01。

	Cycle *string `json:"cycle,omitempty"`
	// 消费日期，格式为YYYY-MM-DD。  说明： 当statistic_type=2时该字段才有值，否则返回null。

	BillDate *string `json:"bill_date,omitempty"`
	// 账单类型。 1：消费-新购2：消费-续订3：消费-变更4：退款-退订5：消费-使用8：消费-自动续订9：调账-补偿14：消费-服务支持计划月末扣费16：调账-扣费18：消费-按月付费20：退款-变更

	BillType *int32 `json:"bill_type,omitempty"`
	// 消费的客户账号ID。 如果是普通客户或者企业子客户查询消费记录，只能查询到客户自己的消费记录，且此处显示的是客户自己的客户ID。如果是企业主查询消费记录，可以查询到企业主以及企业子客户的消费记录，此处为消费的实际客户ID。如果是企业主自己的消费记录，则为企业主ID；如果是某个企业子客户的消费记录，则此处为企业子账号ID。

	CustomerId *string `json:"customer_id,omitempty"`
	// 云服务区编码，例如：“cn-north-1”。具体请参见地区和终端节点对应云服务的“区域”列的值。

	Region *string `json:"region,omitempty"`
	// 云服务区名称，例如：“华北-北京一”。具体请参见地区和终端节点对应云服务的“区域名称”列的值。

	RegionName *string `json:"region_name,omitempty"`
	// 云服务类型编码，例如OBS的云服务类型编码为“hws.service.type.obs”。您可以调用查询云服务类型列表接口获取。

	CloudServiceType *string `json:"cloud_service_type,omitempty"`
	// 资源类型编码，例如ECS的VM为“hws.resource.type.vm”。您可以调用查询资源类型列表接口获取。

	ResourceTypeCode *string `json:"resource_Type_code,omitempty"`
	// 云服务类型名称。例如ECS的云服务类型名称为“弹性云服务器”。

	CloudServiceTypeName *string `json:"cloud_service_type_name,omitempty"`
	// 资源类型名称。例如ECS的资源类型名称为“云主机”。

	ResourceTypeName *string `json:"resource_type_name,omitempty"`
	// 资源实例ID。

	ResInstanceId *string `json:"res_instance_id,omitempty"`
	// 资源名称。客户在创建资源的时候，可以输入资源名称，有些资源也可以在管理资源时，修改资源名称。

	ResourceName *string `json:"resource_name,omitempty"`
	// 资源标签。客户在管理资源的时候，可以设置资源标签。

	ResourceTag *string `json:"resource_tag,omitempty"`
	// SKU编码，在账单中唯一标识一个资源的规格。

	SkuCode *string `json:"sku_code,omitempty"`
	// 企业项目标识（企业项目ID）。 default项目对应ID：0未归集（表示该云服务不支持企业项目管理能力）项目对应ID：-1其余项目对应ID获取方法请参见如何获取企业项目ID。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 企业项目名称。

	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`
	// 计费模式。 1 : 包年/包月3：按需10：预留实例

	ChargeMode *int32 `json:"charge_mode,omitempty"`
	// 客户购买云服务类型的消费金额，包含代金券、现金券，精确到小数点后2位。  说明： consume_amount的值等于cash_amount，credit_amount，coupon_amount，flexipurchase_coupon_amount，stored_card_amount，bonus_amount，debt_amount，adjustment_amount的总和。

	ConsumeAmount *float64 `json:"consume_amount,omitempty"`
	// 现金支付金额。

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
	// 官网价。

	OfficialAmount *float64 `json:"official_amount,omitempty"`
	// 对应官网价折扣金额。

	DiscountAmount *float64 `json:"discount_amount,omitempty"`
	// 金额单位。 1：元

	MeasureId *int32 `json:"measure_id,omitempty"`
	// 周期类型： 19：年20：月24：天25：小时5：一次性

	PeriodType *int32 `json:"period_type,omitempty"`
}

func (o MonthlyBillRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MonthlyBillRes struct{}"
	}

	return strings.Join([]string{"MonthlyBillRes", string(data)}, " ")
}
