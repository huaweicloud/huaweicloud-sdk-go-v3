package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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
