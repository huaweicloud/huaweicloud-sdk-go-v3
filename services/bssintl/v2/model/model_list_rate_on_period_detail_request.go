package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListRateOnPeriodDetailRequest struct {
	Body *RateOnPeriodReq `json:"body,omitempty" xml:"body"`
}

func (o ListRateOnPeriodDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRateOnPeriodDetailRequest struct{}"
	}

	return strings.Join([]string{"ListRateOnPeriodDetailRequest", string(data)}, " ")
}
