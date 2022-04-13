package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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
