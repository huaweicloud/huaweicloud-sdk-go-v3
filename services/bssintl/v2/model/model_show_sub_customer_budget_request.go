package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowSubCustomerBudgetRequest struct {
	// 客户账号ID。您可以调用查询客户列表接口获取customer_id。

	CustomerId string `json:"customer_id"`
}

func (o ShowSubCustomerBudgetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSubCustomerBudgetRequest struct{}"
	}

	return strings.Join([]string{"ShowSubCustomerBudgetRequest", string(data)}, " ")
}
