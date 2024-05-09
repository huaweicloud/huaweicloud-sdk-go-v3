package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"
	"strings"
)

type BudgetInfo struct {

	// |参数名称：客户账号ID||参数的约束及描述：必填参数，范围限制:1-64|
	CustomerId *string `json:"customer_id,omitempty"`

	// |参数名称：初始预算金额。| |参数的约束及描述：非必填，初始预算金额。|
	BudgetAmount *decimal.Decimal `json:"budget_amount,omitempty"`

	// |参数名称：已经使用的预算。| |参数的约束及描述：已经使用的预算。非必填，该预算存在一定的时延和误差。|
	UsedAmount *decimal.Decimal `json:"used_amount,omitempty"`

	// |参数名称：预算模式| |参数的约束及描述：MONTHLY 月度预算 PACKAGE 一次性预算|
	BudgetType *BudgetInfoBudgetType `json:"budget_type,omitempty"`
}

func (o BudgetInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BudgetInfo struct{}"
	}

	return strings.Join([]string{"BudgetInfo", string(data)}, " ")
}

type BudgetInfoBudgetType struct {
	value string
}

type BudgetInfoBudgetTypeEnum struct {
	MONTHLY BudgetInfoBudgetType
	PACKAGE BudgetInfoBudgetType
}

func GetBudgetInfoBudgetTypeEnum() BudgetInfoBudgetTypeEnum {
	return BudgetInfoBudgetTypeEnum{
		MONTHLY: BudgetInfoBudgetType{
			value: "MONTHLY",
		},
		PACKAGE: BudgetInfoBudgetType{
			value: "PACKAGE",
		},
	}
}

func (c BudgetInfoBudgetType) Value() string {
	return c.value
}

func (c BudgetInfoBudgetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BudgetInfoBudgetType) UnmarshalJSON(b []byte) error {
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
