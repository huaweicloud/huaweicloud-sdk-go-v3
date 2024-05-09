package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubCustomerBudgetResponse Response Object
type ListSubCustomerBudgetResponse struct {

	// |参数名称：金额单位。||参数的约束及描述：非必填参数|
	MeasureId *int32 `json:"measure_id,omitempty"`

	// |参数名称：货币，CNY：人民币，USD：美元||参数的约束及描述：非必填参数|
	Currency *string `json:"currency,omitempty"`

	// |参数名称：客户预算信息||参数的约束及描述：必填参数|
	BudgetInfos    *[]BudgetInfo `json:"budget_infos,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListSubCustomerBudgetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubCustomerBudgetResponse struct{}"
	}

	return strings.Join([]string{"ListSubCustomerBudgetResponse", string(data)}, " ")
}
