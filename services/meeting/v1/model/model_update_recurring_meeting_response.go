package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRecurringMeetingResponse Response Object
type UpdateRecurringMeetingResponse struct {

	// 会议信息列表。
	Body           *[]ConferenceInfo `json:"body,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o UpdateRecurringMeetingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRecurringMeetingResponse struct{}"
	}

	return strings.Join([]string{"UpdateRecurringMeetingResponse", string(data)}, " ")
}
