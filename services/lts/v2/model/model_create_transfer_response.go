package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTransferResponse Response Object
type CreateTransferResponse struct {

	// 日志组ID
	LogGroupId *string `json:"log_group_id,omitempty"`

	// 日志组名称
	LogGroupName *string `json:"log_group_name,omitempty"`

	// 日志流ID集合
	LogStreams *[]CreateTransferResponseBodyLogStreams `json:"log_streams,omitempty"`

	// 日志转储ID
	LogTransferId *string `json:"log_transfer_id,omitempty"`

	LogTransferInfo *CreateTransferResponseBodyLogTransferInfo `json:"log_transfer_info,omitempty"`
	HttpStatusCode  int                                        `json:"-"`
}

func (o CreateTransferResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTransferResponse struct{}"
	}

	return strings.Join([]string{"CreateTransferResponse", string(data)}, " ")
}
