package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetSensitiveResultSwitchResponse Response Object
type SetSensitiveResultSwitchResponse struct {

	// 状态  - SUCCESS：成功  - FAILED：失败
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetSensitiveResultSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetSensitiveResultSwitchResponse struct{}"
	}

	return strings.Join([]string{"SetSensitiveResultSwitchResponse", string(data)}, " ")
}
