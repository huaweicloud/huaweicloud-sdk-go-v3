package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateTransferResponse struct {

	// 日志组ID
	LogGroupId *string `json:"log_group_id,omitempty" xml:"log_group_id"`

	// 日志组名称
	LogGroupName *string `json:"log_group_name,omitempty" xml:"log_group_name"`

	// 日志流ID集合
	LogStreams *[]CreateTransferResponseBodyLogStreams `json:"log_streams,omitempty" xml:"log_streams"`

	// 日志转储ID
	LogTransferId *string `json:"log_transfer_id,omitempty" xml:"log_transfer_id"`

	LogTransferInfo *CreateTransferResponseBodyLogTransferInfo `json:"log_transfer_info,omitempty" xml:"log_transfer_info"`
	HttpStatusCode  int                                        `json:"-"`
}

func (o UpdateTransferResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTransferResponse struct{}"
	}

	return strings.Join([]string{"UpdateTransferResponse", string(data)}, " ")
}
