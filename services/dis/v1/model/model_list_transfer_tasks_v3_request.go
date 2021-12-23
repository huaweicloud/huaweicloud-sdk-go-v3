package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListTransferTasksV3Request struct {
	// 需要查询的通道名称。

	StreamName string `json:"stream_name"`
}

func (o ListTransferTasksV3Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTransferTasksV3Request struct{}"
	}

	return strings.Join([]string{"ListTransferTasksV3Request", string(data)}, " ")
}
