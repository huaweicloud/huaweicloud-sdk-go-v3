package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMonthUsagesRequest Request Object
type ShowMonthUsagesRequest struct {
	Body *ShowMonthUsageReq `json:"body,omitempty"`
}

func (o ShowMonthUsagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMonthUsagesRequest struct{}"
	}

	return strings.Join([]string{"ShowMonthUsagesRequest", string(data)}, " ")
}
