package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelScheduleTaskResponse Response Object
type CancelScheduleTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CancelScheduleTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelScheduleTaskResponse struct{}"
	}

	return strings.Join([]string{"CancelScheduleTaskResponse", string(data)}, " ")
}
