package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAlarmResponse Response Object
type ShowAlarmResponse struct {
	Data           *AlarmInfoDto `json:"data,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowAlarmResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlarmResponse struct{}"
	}

	return strings.Join([]string{"ShowAlarmResponse", string(data)}, " ")
}
