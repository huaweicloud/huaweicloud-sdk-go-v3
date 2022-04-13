package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建OBS转储，DIS转储，DMS转储
type CreateTransferRequestBody struct {
	// 日志组ID

	LogGroupId string `json:"log_group_id"`
	// 日志流ID集合

	LogStreams []CreateTransferRequestBodyLogStreams `json:"log_streams"`

	LogTransferInfo *CreateTransferRequestBodyLogTransferInfo `json:"log_transfer_info"`
}

func (o CreateTransferRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTransferRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTransferRequestBody", string(data)}, " ")
}
