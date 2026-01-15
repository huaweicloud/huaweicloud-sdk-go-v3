package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetSensitiveSwitchResponse Response Object
type SetSensitiveSwitchResponse struct {

	// 状态  - SUCCESS：成功  - FAILED：失败
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetSensitiveSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetSensitiveSwitchResponse struct{}"
	}

	return strings.Join([]string{"SetSensitiveSwitchResponse", string(data)}, " ")
}
