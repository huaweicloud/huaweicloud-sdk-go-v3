package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowMonthUsagesRequest struct {
	Body *ShowMonthUsageReq `json:"body,omitempty" xml:"body"`
}

func (o ShowMonthUsagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMonthUsagesRequest struct{}"
	}

	return strings.Join([]string{"ShowMonthUsagesRequest", string(data)}, " ")
}
