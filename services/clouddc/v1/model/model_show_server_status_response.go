package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServerStatusResponse Response Object
type ShowServerStatusResponse struct {
	ServerStatus *ServerStatus `json:"server_status,omitempty"`

	ServerPowerStatus *ServerPowerStatus `json:"server_power_status,omitempty"`
	HttpStatusCode    int                `json:"-"`
}

func (o ShowServerStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowServerStatusResponse", string(data)}, " ")
}
