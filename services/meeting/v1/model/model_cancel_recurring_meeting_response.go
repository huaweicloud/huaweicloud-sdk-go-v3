package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelRecurringMeetingResponse Response Object
type CancelRecurringMeetingResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CancelRecurringMeetingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelRecurringMeetingResponse struct{}"
	}

	return strings.Join([]string{"CancelRecurringMeetingResponse", string(data)}, " ")
}
