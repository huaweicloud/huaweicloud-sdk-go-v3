package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateTransferResponseBody struct {
	// 日志组ID

	LogGroupId string `json:"log_group_id"`
	// 日志组名称

	LogGroupName string `json:"log_group_name"`
	// 日志流ID集合

	LogStreams []CreateTransferResponseBodyLogStreams `json:"log_streams"`
	// 日志转储ID

	LogTransferId string `json:"log_transfer_id"`

	LogTransferInfo *CreateTransferResponseBodyLogTransferInfo `json:"log_transfer_info"`
}

func (o CreateTransferResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTransferResponseBody struct{}"
	}

	return strings.Join([]string{"CreateTransferResponseBody", string(data)}, " ")
}
