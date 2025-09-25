package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSerialConsoleOptionsRequest Request Object
type UpdateSerialConsoleOptionsRequest struct {

	// 云服务器ID。
	ServerId string `json:"server_id"`

	Body *UpdateSerialConsoleOptionsRequestBody `json:"body,omitempty"`
}

func (o UpdateSerialConsoleOptionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSerialConsoleOptionsRequest struct{}"
	}

	return strings.Join([]string{"UpdateSerialConsoleOptionsRequest", string(data)}, " ")
}
