package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteScheduleResponse Response Object
type DeleteScheduleResponse struct {
	Data           *ScheduleRespData `json:"data,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o DeleteScheduleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScheduleResponse struct{}"
	}

	return strings.Join([]string{"DeleteScheduleResponse", string(data)}, " ")
}
