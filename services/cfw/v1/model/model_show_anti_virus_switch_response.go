package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAntiVirusSwitchResponse Response Object
type ShowAntiVirusSwitchResponse struct {
	Data           *AntiVirusVo `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowAntiVirusSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAntiVirusSwitchResponse struct{}"
	}

	return strings.Join([]string{"ShowAntiVirusSwitchResponse", string(data)}, " ")
}
