package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopMeetingResponse Response Object
type StopMeetingResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopMeetingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopMeetingResponse struct{}"
	}

	return strings.Join([]string{"StopMeetingResponse", string(data)}, " ")
}
