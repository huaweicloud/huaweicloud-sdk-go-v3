package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTransfersResponse Response Object
type ListTransfersResponse struct {

	// 查询日志转储信息数组
	LogTransfers   *[]CreateTransferResponseBody `json:"log_transfers,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListTransfersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTransfersResponse struct{}"
	}

	return strings.Join([]string{"ListTransfersResponse", string(data)}, " ")
}
