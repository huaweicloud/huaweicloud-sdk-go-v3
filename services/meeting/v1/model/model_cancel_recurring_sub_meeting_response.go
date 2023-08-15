package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelRecurringSubMeetingResponse Response Object
type CancelRecurringSubMeetingResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CancelRecurringSubMeetingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelRecurringSubMeetingResponse struct{}"
	}

	return strings.Join([]string{"CancelRecurringSubMeetingResponse", string(data)}, " ")
}
