package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteScheduleTaskResponse Response Object
type DeleteScheduleTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteScheduleTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScheduleTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteScheduleTaskResponse", string(data)}, " ")
}
