package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetAlarmTopicConfigInfoResponse Response Object
type SetAlarmTopicConfigInfoResponse struct {

	// 状态  - success：成功  - fail：失败
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetAlarmTopicConfigInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAlarmTopicConfigInfoResponse struct{}"
	}

	return strings.Join([]string{"SetAlarmTopicConfigInfoResponse", string(data)}, " ")
}
