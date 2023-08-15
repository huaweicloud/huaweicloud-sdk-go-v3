package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMeetingResponse Response Object
type CreateMeetingResponse struct {

	// 会议信息列表。
	Body           *[]ConferenceInfo `json:"body,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o CreateMeetingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMeetingResponse struct{}"
	}

	return strings.Join([]string{"CreateMeetingResponse", string(data)}, " ")
}
