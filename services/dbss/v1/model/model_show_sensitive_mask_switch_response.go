package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSensitiveMaskSwitchResponse Response Object
type ShowSensitiveMaskSwitchResponse struct {

	// 状态  - ON：开启  - OFF：关闭
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSensitiveMaskSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSensitiveMaskSwitchResponse struct{}"
	}

	return strings.Join([]string{"ShowSensitiveMaskSwitchResponse", string(data)}, " ")
}
