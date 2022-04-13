package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StartMeetingRequest struct {
	Body *StartRequest `json:"body,omitempty"`
}

func (o StartMeetingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartMeetingRequest struct{}"
	}

	return strings.Join([]string{"StartMeetingRequest", string(data)}, " ")
}
