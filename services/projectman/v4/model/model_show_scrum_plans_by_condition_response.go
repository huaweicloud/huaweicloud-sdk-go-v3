package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowScrumPlansByConditionResponse Response Object
type ShowScrumPlansByConditionResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowScrumPlansByConditionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScrumPlansByConditionResponse struct{}"
	}

	return strings.Join([]string{"ShowScrumPlansByConditionResponse", string(data)}, " ")
}
