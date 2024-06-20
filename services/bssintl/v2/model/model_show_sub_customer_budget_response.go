package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"
	"strings"
)

// ShowSubCustomerBudgetResponse Response Object
type ShowSubCustomerBudgetResponse struct {

	// 初始预算金额。
	BudgetAmount *decimal.Decimal `json:"budget_amount,omitempty"`

	// 已经使用的预算。该预算存在一定的时延和误差。
	UsedAmount *decimal.Decimal `json:"used_amount,omitempty"`

	// 金额单位。 1：元
	MeasureId *int32 `json:"measure_id,omitempty"`

	// 币种。 USD：美金
	Currency *string `json:"currency,omitempty"`

	// |参数名称：预算模式| |参数的约束及描述：MONTHLY 月度预算 PACKAGE 一次性预算|
	BudgetType     *ShowSubCustomerBudgetResponseBudgetType `json:"budget_type,omitempty"`
	HttpStatusCode int                                      `json:"-"`
}

func (o ShowSubCustomerBudgetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSubCustomerBudgetResponse struct{}"
	}

	return strings.Join([]string{"ShowSubCustomerBudgetResponse", string(data)}, " ")
}

type ShowSubCustomerBudgetResponseBudgetType struct {
	value string
}

type ShowSubCustomerBudgetResponseBudgetTypeEnum struct {
	MONTHLY ShowSubCustomerBudgetResponseBudgetType
	PACKAGE ShowSubCustomerBudgetResponseBudgetType
}

func GetShowSubCustomerBudgetResponseBudgetTypeEnum() ShowSubCustomerBudgetResponseBudgetTypeEnum {
	return ShowSubCustomerBudgetResponseBudgetTypeEnum{
		MONTHLY: ShowSubCustomerBudgetResponseBudgetType{
			value: "MONTHLY",
		},
		PACKAGE: ShowSubCustomerBudgetResponseBudgetType{
			value: "PACKAGE",
		},
	}
}

func (c ShowSubCustomerBudgetResponseBudgetType) Value() string {
	return c.value
}

func (c ShowSubCustomerBudgetResponseBudgetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSubCustomerBudgetResponseBudgetType) UnmarshalJSON(b []byte) error {
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
