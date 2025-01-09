package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteScheduleTaskResponse Response Object
type BatchDeleteScheduleTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteScheduleTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteScheduleTaskResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteScheduleTaskResponse", string(data)}, " ")
}
