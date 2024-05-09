package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubCustomerBudgetRequest Request Object
type ListSubCustomerBudgetRequest struct {
	Body *ListSubCustomerBudgetReq `json:"body,omitempty"`
}

func (o ListSubCustomerBudgetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubCustomerBudgetRequest struct{}"
	}

	return strings.Join([]string{"ListSubCustomerBudgetRequest", string(data)}, " ")
}
