package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServerRemoteConsoleReq This is a auto create Body Object
type ShowServerRemoteConsoleReq struct {
	RemoteConsole *RemoteConsole `json:"remote_console"`
}

func (o ShowServerRemoteConsoleReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerRemoteConsoleReq struct{}"
	}

	return strings.Join([]string{"ShowServerRemoteConsoleReq", string(data)}, " ")
}
