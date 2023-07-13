package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScaleOut 扩容集群详情
type ScaleOut struct {

	// 扩容节点数。
	Count int32 `json:"count"`
}

func (o ScaleOut) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScaleOut struct{}"
	}

	return strings.Join([]string{"ScaleOut", string(data)}, " ")
}
