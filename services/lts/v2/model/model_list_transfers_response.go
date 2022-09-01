package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListTransfersResponse struct {

	// 查询日志转储信息数组
	LogTransfers   *[]CreateTransferResponseBody `json:"log_transfers,omitempty" xml:"log_transfers"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListTransfersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTransfersResponse struct{}"
	}

	return strings.Join([]string{"ListTransfersResponse", string(data)}, " ")
}
