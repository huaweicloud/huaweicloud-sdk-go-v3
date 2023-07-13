package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAlarmResponse Response Object
type UpdateAlarmResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAlarmResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmResponse struct{}"
	}

	return strings.Join([]string{"UpdateAlarmResponse", string(data)}, " ")
}
