package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteStreamV3Request struct {
	// 需要删除的通道名称。

	StreamName string `json:"stream_name"`
}

func (o DeleteStreamV3Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStreamV3Request struct{}"
	}

	return strings.Join([]string{"DeleteStreamV3Request", string(data)}, " ")
}
