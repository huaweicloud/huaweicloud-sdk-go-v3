package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProlongMeetingResponse Response Object
type ProlongMeetingResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ProlongMeetingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProlongMeetingResponse struct{}"
	}

	return strings.Join([]string{"ProlongMeetingResponse", string(data)}, " ")
}
