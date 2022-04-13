package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListCustomersBalancesDetailRequest struct {
	Body *QueryCustomersBalancesReq `json:"body,omitempty"`
}

func (o ListCustomersBalancesDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomersBalancesDetailRequest struct{}"
	}

	return strings.Join([]string{"ListCustomersBalancesDetailRequest", string(data)}, " ")
}
