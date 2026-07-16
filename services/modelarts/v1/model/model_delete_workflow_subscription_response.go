package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWorkflowSubscriptionResponse Response Object
type DeleteWorkflowSubscriptionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteWorkflowSubscriptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWorkflowSubscriptionResponse struct{}"
	}

	return strings.Join([]string{"DeleteWorkflowSubscriptionResponse", string(data)}, " ")
}
