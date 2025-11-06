package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HandlerAlarmResponse Response Object
type HandlerAlarmResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o HandlerAlarmResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HandlerAlarmResponse struct{}"
	}

	return strings.Join([]string{"HandlerAlarmResponse", string(data)}, " ")
}
