package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServerRemoteConsoleRequestBody This is a auto create Body Object
type ShowServerRemoteConsoleRequestBody struct {
	RemoteConsole *GetServerRemoteConsoleOption `json:"remote_console"`
}

func (o ShowServerRemoteConsoleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerRemoteConsoleRequestBody struct{}"
	}

	return strings.Join([]string{"ShowServerRemoteConsoleRequestBody", string(data)}, " ")
}
