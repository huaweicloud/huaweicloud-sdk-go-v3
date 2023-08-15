package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServerRemoteConsoleRequest Request Object
type ShowServerRemoteConsoleRequest struct {
	ServerId string `json:"server_id"`

	Body *ShowServerRemoteConsoleReq `json:"body,omitempty"`
}

func (o ShowServerRemoteConsoleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerRemoteConsoleRequest struct{}"
	}

	return strings.Join([]string{"ShowServerRemoteConsoleRequest", string(data)}, " ")
}
