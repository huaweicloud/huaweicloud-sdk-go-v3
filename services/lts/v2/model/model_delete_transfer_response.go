package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteTransferResponse struct {
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

func (o DeleteTransferResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTransferResponse struct{}"
	}

	return strings.Join([]string{"DeleteTransferResponse", string(data)}, " ")
}
