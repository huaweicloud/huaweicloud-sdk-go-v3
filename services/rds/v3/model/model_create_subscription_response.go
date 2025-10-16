package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubscriptionResponse Response Object
type CreateSubscriptionResponse struct {

	// 任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSubscriptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubscriptionResponse struct{}"
	}

	return strings.Join([]string{"CreateSubscriptionResponse", string(data)}, " ")
}
