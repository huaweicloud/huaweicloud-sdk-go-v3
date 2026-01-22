package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateScheduleResponse Response Object
type CreateScheduleResponse struct {
	Data           *ScheduleRespData `json:"data,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o CreateScheduleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScheduleResponse struct{}"
	}

	return strings.Join([]string{"CreateScheduleResponse", string(data)}, " ")
}
