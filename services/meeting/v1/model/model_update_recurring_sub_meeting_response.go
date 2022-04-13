package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateRecurringSubMeetingResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateRecurringSubMeetingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRecurringSubMeetingResponse struct{}"
	}

	return strings.Join([]string{"UpdateRecurringSubMeetingResponse", string(data)}, " ")
}
