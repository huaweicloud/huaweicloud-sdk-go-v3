package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LockMeetingResponse Response Object
type LockMeetingResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o LockMeetingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LockMeetingResponse struct{}"
	}

	return strings.Join([]string{"LockMeetingResponse", string(data)}, " ")
}
