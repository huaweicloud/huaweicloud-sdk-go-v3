package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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
