package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPolicyEventsRequest Request Object
type ListPolicyEventsRequest struct {

	// 策略id
	Id string `json:"id"`
}

func (o ListPolicyEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyEventsRequest struct{}"
	}

	return strings.Join([]string{"ListPolicyEventsRequest", string(data)}, " ")
}
