package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKillProcessTaskSwitchResponse Response Object
type ShowKillProcessTaskSwitchResponse struct {

	// 开关状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowKillProcessTaskSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKillProcessTaskSwitchResponse struct{}"
	}

	return strings.Join([]string{"ShowKillProcessTaskSwitchResponse", string(data)}, " ")
}
