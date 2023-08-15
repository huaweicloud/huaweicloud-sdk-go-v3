package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSubCustomerBudgetRequest Request Object
type UpdateSubCustomerBudgetRequest struct {
	Body *ModSubCustomerBudgetReq `json:"body,omitempty"`
}

func (o UpdateSubCustomerBudgetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubCustomerBudgetRequest struct{}"
	}

	return strings.Join([]string{"UpdateSubCustomerBudgetRequest", string(data)}, " ")
}
