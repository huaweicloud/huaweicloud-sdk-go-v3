package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

type BillSumInfoV2 struct {

	// 客户账号ID。
	CustomerId *string `json:"customer_id,omitempty"`

	// 云服务类型编码，例如OBS的云服务类型编码为“hws.service.type.obs”。
	CloudServiceType *string `json:"cloud_service_type,omitempty"`

	// 云服务类型名称。例如ECS的云服务类型名称为“弹性云服务器”。
	CloudServiceTypeName *string `json:"cloud_service_type_name,omitempty"`

	// 账单类型。 0：正常1：退订2：华为核销
	BillType *string `json:"bill_type,omitempty"`

	// 计费模式。 1：包年/包月3: 按需
	ChargeMode *string `json:"charge_mode,omitempty"`

	// 金额。 对于billType=1或者2的账单，该金额为负值。
	Amount *decimal.Decimal `json:"amount,omitempty"`

	// 欠费金额，指从客户账户扣费的时候，客户账户金额不足，欠费的金额，华为核销或者退订的时候没有该字段。
	DebtAmount *decimal.Decimal `json:"debt_amount,omitempty"`

	// 核销欠款，华为核销或者退订的时候没有该字段。
	AdjustmentAmount *decimal.Decimal `json:"adjustment_amount,omitempty"`

	// 折扣金额，华为核销或者退订的时候没有该字段。
	DiscountAmount *decimal.Decimal `json:"discount_amount,omitempty"`

	// 金额单位。 1：元
	MeasureId *int32 `json:"measure_id,omitempty"`

	// 按不同账户消费类型和付费方式区分的支付总金额。 具体请参见表3。
	AccountDetails *[]BalanceTypeDeductSumV2 `json:"account_details,omitempty"`

	// 资源类型编码，例如ECS的VM为“hws.resource.type.vm”。
	ResourceTypeCode *string `json:"resource_type_code,omitempty"`

	// 资源类型名称。例如ECS的资源类型名称为“云主机”。
	ResourceTypeName *string `json:"resource_type_name,omitempty"`
}

func (o BillSumInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BillSumInfoV2 struct{}"
	}

	return strings.Join([]string{"BillSumInfoV2", string(data)}, " ")
}
