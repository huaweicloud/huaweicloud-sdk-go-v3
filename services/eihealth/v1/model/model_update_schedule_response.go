package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScheduleResponse Response Object
type UpdateScheduleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateScheduleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScheduleResponse struct{}"
	}

	return strings.Join([]string{"UpdateScheduleResponse", string(data)}, " ")
}
