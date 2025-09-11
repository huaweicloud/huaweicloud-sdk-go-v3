package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSensitiveResultSwitchResponse Response Object
type ShowSensitiveResultSwitchResponse struct {

	// 状态  - ON：开启  - OFF：关闭
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSensitiveResultSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSensitiveResultSwitchResponse struct{}"
	}

	return strings.Join([]string{"ShowSensitiveResultSwitchResponse", string(data)}, " ")
}
