package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CancelMeetingResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CancelMeetingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelMeetingResponse struct{}"
	}

	return strings.Join([]string{"CancelMeetingResponse", string(data)}, " ")
}
