package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HandleAlarmResponse Response Object
type HandleAlarmResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o HandleAlarmResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HandleAlarmResponse struct{}"
	}

	return strings.Join([]string{"HandleAlarmResponse", string(data)}, " ")
}
