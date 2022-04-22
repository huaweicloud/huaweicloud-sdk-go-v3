package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateRecurringMeetingResponse struct {

	// 会议信息列表
	Body           *[]ConferenceInfo `json:"body,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o CreateRecurringMeetingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRecurringMeetingResponse struct{}"
	}

	return strings.Join([]string{"CreateRecurringMeetingResponse", string(data)}, " ")
}
