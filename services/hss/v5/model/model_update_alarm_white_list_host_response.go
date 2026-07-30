package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAlarmWhiteListHostResponse Response Object
type UpdateAlarmWhiteListHostResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAlarmWhiteListHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmWhiteListHostResponse struct{}"
	}

	return strings.Join([]string{"UpdateAlarmWhiteListHostResponse", string(data)}, " ")
}
