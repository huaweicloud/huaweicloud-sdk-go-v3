package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateTransferRequestBody struct {

	// 日志转储ID
	LogTransferId string `json:"log_transfer_id"`

	// 日志组ID
	LogGroupId *string `json:"log_group_id,omitempty"`

	LogTransferInfo *UpdateTransferRequestBodyLogTransferInfo `json:"log_transfer_info"`

	// 日志流信息
	LogStreams *[]LogStreams `json:"log_streams,omitempty"`
}

func (o UpdateTransferRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTransferRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateTransferRequestBody", string(data)}, " ")
}
