package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteAlarmResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAlarmResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlarmResponse struct{}"
	}

	return strings.Join([]string{"DeleteAlarmResponse", string(data)}, " ")
}
