package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSerialConsoleActionsOption
type ShowSerialConsoleActionsOption struct {

	// 串口登录的URL
	Url *string `json:"url,omitempty"`
}

func (o ShowSerialConsoleActionsOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSerialConsoleActionsOption struct{}"
	}

	return strings.Join([]string{"ShowSerialConsoleActionsOption", string(data)}, " ")
}
