package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWorkflowScheduleIdResponse Response Object
type DeleteWorkflowScheduleIdResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteWorkflowScheduleIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWorkflowScheduleIdResponse struct{}"
	}

	return strings.Join([]string{"DeleteWorkflowScheduleIdResponse", string(data)}, " ")
}
