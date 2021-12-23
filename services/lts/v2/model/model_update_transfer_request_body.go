package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateTransferRequestBody struct {
	// 日志转储ID

	LogTransferId string `json:"log_transfer_id"`

	LogTransferInfo *UpdateTransferRequestBodyLogTransferInfo `json:"log_transfer_info"`
}

func (o UpdateTransferRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTransferRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateTransferRequestBody", string(data)}, " ")
}
