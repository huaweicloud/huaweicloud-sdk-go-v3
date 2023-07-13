package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRenewRateOnPeriodRequest Request Object
type ListRenewRateOnPeriodRequest struct {
	Body *ListRenewRateOnPeriodReq `json:"body,omitempty"`
}

func (o ListRenewRateOnPeriodRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRenewRateOnPeriodRequest struct{}"
	}

	return strings.Join([]string{"ListRenewRateOnPeriodRequest", string(data)}, " ")
}
