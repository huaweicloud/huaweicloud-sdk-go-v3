package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteScheduleRecordTasksResponse Response Object
type DeleteScheduleRecordTasksResponse struct {
	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteScheduleRecordTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScheduleRecordTasksResponse struct{}"
	}

	return strings.Join([]string{"DeleteScheduleRecordTasksResponse", string(data)}, " ")
}
