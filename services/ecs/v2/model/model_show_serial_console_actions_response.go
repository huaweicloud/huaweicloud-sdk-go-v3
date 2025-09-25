package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSerialConsoleActionsResponse Response Object
type ShowSerialConsoleActionsResponse struct {
	SerialConsole  *ShowSerialConsoleActionsOption `json:"serial_console,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ShowSerialConsoleActionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSerialConsoleActionsResponse struct{}"
	}

	return strings.Join([]string{"ShowSerialConsoleActionsResponse", string(data)}, " ")
}
