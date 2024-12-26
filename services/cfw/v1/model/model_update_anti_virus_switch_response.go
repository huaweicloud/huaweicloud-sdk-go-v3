package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAntiVirusSwitchResponse Response Object
type UpdateAntiVirusSwitchResponse struct {
	Data           *ResponseData `json:"data,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o UpdateAntiVirusSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAntiVirusSwitchResponse struct{}"
	}

	return strings.Join([]string{"UpdateAntiVirusSwitchResponse", string(data)}, " ")
}
