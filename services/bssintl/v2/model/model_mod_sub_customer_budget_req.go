package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ModSubCustomerBudgetReq struct {

	// 客户账号ID。您可以调用查询客户列表接口获取customer_id。
	CustomerId string `json:"customer_id"`

	// 调整的目标金额。 单位：元。精确至小数点后2位。
	BudgetAmount float64 `json:"budget_amount"`

	// 是否在设置客户预算的同时解除账号冻结： 0：否1：是 默认值为0。
	CancelPartnerFrozen *string `json:"cancel_partner_frozen,omitempty"`

	// 云经销商ID。获取方法请参见查询云经销商列表。如果需要查询云经销商的子客户列表，必须携带该字段。除此之外，此参数不做处理。
	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`

	// |参数名称：预算模式| |参数的约束及描述：MONTHLY 月度预算 PACKAGE 一次性预算 ，此参数不携带或携带值为null时，默认值为MONTHLY。|
	BudgetType *ModSubCustomerBudgetReqBudgetType `json:"budget_type,omitempty"`
}

func (o ModSubCustomerBudgetReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModSubCustomerBudgetReq struct{}"
	}

	return strings.Join([]string{"ModSubCustomerBudgetReq", string(data)}, " ")
}

type ModSubCustomerBudgetReqBudgetType struct {
	value string
}

type ModSubCustomerBudgetReqBudgetTypeEnum struct {
	MONTHLY ModSubCustomerBudgetReqBudgetType
	PACKAGE ModSubCustomerBudgetReqBudgetType
}

func GetModSubCustomerBudgetReqBudgetTypeEnum() ModSubCustomerBudgetReqBudgetTypeEnum {
	return ModSubCustomerBudgetReqBudgetTypeEnum{
		MONTHLY: ModSubCustomerBudgetReqBudgetType{
			value: "MONTHLY",
		},
		PACKAGE: ModSubCustomerBudgetReqBudgetType{
			value: "PACKAGE",
		},
	}
}

func (c ModSubCustomerBudgetReqBudgetType) Value() string {
	return c.value
}

func (c ModSubCustomerBudgetReqBudgetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModSubCustomerBudgetReqBudgetType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
