package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAlarmResponse Response Object
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
