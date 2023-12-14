package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTransferTasksRequest Request Object
type ListTransferTasksRequest struct {

	// 需要查询的通道名称。
	StreamName string `json:"stream_name"`
}

func (o ListTransferTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTransferTasksRequest struct{}"
	}

	return strings.Join([]string{"ListTransferTasksRequest", string(data)}, " ")
}
