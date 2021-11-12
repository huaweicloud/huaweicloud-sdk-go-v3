package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteStreamRequest struct {
	// 需要删除的通道名称。

	StreamName string `json:"stream_name"`
}

func (o DeleteStreamRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStreamRequest struct{}"
	}

	return strings.Join([]string{"DeleteStreamRequest", string(data)}, " ")
}
