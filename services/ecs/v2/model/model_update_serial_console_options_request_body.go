package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSerialConsoleOptionsRequestBody This is a auto create Body Object
type UpdateSerialConsoleOptionsRequestBody struct {
	SerialConsoleOptions *UpdateSerialConsoleOptionsOption `json:"serial_console_options,omitempty"`
}

func (o UpdateSerialConsoleOptionsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSerialConsoleOptionsRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSerialConsoleOptionsRequestBody", string(data)}, " ")
}
