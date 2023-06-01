package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

type BillSumRecordInfoV2 struct {

	// 消费汇总数据所在账期，东八区时间，格式：YYYY-MM。
	BillCycle *string `json:"bill_cycle,omitempty"`

	// 资源类型编码，例如ECS的VM为“hws.resource.type.vm”。您可以调用查询资源类型列表接口获取。
	ResourceTypeCode *string `json:"resource_type_code,omitempty"`

	// 云服务类型编码，例如OBS的云服务类型编码为“hws.service.type.obs”。您可以调用查询云服务类型列表接口获取。
	ServiceTypeCode *string `json:"service_type_code,omitempty"`

	// 云服务类型名称。例如ECS的云服务类型名称为“弹性云服务器”。
	ServiceTypeName *string `json:"service_type_name,omitempty"`

	// 资源类型名称。例如ECS的资源类型名称为“云主机”。
	ResourceTypeName *string `json:"resource_type_name,omitempty"`

	// 计费模式。 1：包年/包月3：按需10：预留实例
	ChargingMode *int32 `json:"charging_mode,omitempty"`

	// 官网价。
	OfficialAmount *decimal.Decimal `json:"official_amount,omitempty"`

	// 折扣金额。
	OfficialDiscountAmount *decimal.Decimal `json:"official_discount_amount,omitempty"`

	// 抹零金额。
	TruncatedAmount *decimal.Decimal `json:"truncated_amount,omitempty"`

	// 应付金额。 应付金额=官网价-折扣金额-抹零金额
	ConsumeAmount *decimal.Decimal `json:"consume_amount,omitempty"`

	// 代金券金额。
	CouponAmount *decimal.Decimal `json:"coupon_amount,omitempty"`

	// 现金券金额，预留。
	FlexipurchaseCouponAmount *decimal.Decimal `json:"flexipurchase_coupon_amount,omitempty"`

	// 储值卡金额，预留。
	StoredValueCardAmount *decimal.Decimal `json:"stored_value_card_amount,omitempty"`

	// 欠费金额。即伙伴从客户账户扣费时，客户账户金额不足，欠费的金额。
	DebtAmount *decimal.Decimal `json:"debt_amount,omitempty"`

	// 欠费核销金额。
	WriteoffAmount *decimal.Decimal `json:"writeoff_amount,omitempty"`

	// 现金账户金额。
	CashAmount *decimal.Decimal `json:"cash_amount,omitempty"`

	// 信用账户金额。
	CreditAmount *decimal.Decimal `json:"credit_amount,omitempty"`

	// 金额单位。 1：元
	MeasureId *int32 `json:"measure_id,omitempty"`

	// 账单类型。 1：消费2：退款3：调账
	BillType *int32 `json:"bill_type,omitempty"`

	// 消费的客户账号ID。 如果是普通客户或者企业子客户查询消费记录，只能查询到客户自己的消费记录，且此处显示的是客户自己的客户ID。如果是企业主查询消费记录，可以查询到企业主以及企业子客户的消费记录，此处为消费的实际客户ID。如果是企业主自己的消费记录，则为企业主ID；如果是某个企业子客户的消费记录，则此处为企业子账号ID。
	CustomerId *string `json:"customer_id,omitempty"`
}

func (o BillSumRecordInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BillSumRecordInfoV2 struct{}"
	}

	return strings.Join([]string{"BillSumRecordInfoV2", string(data)}, " ")
}
