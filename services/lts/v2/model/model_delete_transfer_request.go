package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteTransferRequest struct {

	// 日志转储ID。获取ID有3种方式： 1. 调用查询日志转储接口，返回值有日志转储ID  2. 调用新增日志转储接口，返回值有日志转储ID 3. 调用删除日志转储接口，返回值有日志转储ID
	LogTransferId string `json:"log_transfer_id" xml:"log_transfer_id"`
}

func (o DeleteTransferRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTransferRequest struct{}"
	}

	return strings.Join([]string{"DeleteTransferRequest", string(data)}, " ")
}
