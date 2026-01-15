package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetAlarmTopicConfigInfoNewResponse Response Object
type SetAlarmTopicConfigInfoNewResponse struct {

	// 状态  - success：成功  - fail：失败
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetAlarmTopicConfigInfoNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAlarmTopicConfigInfoNewResponse struct{}"
	}

	return strings.Join([]string{"SetAlarmTopicConfigInfoNewResponse", string(data)}, " ")
}
