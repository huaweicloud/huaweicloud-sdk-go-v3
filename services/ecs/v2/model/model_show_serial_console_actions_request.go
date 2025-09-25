package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSerialConsoleActionsRequest Request Object
type ShowSerialConsoleActionsRequest struct {

	// 云服务器ID。
	ServerId string `json:"server_id"`
}

func (o ShowSerialConsoleActionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSerialConsoleActionsRequest struct{}"
	}

	return strings.Join([]string{"ShowSerialConsoleActionsRequest", string(data)}, " ")
}
