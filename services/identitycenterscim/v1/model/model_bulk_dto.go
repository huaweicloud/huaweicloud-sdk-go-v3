package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BulkDto struct {

	// 一个布尔值，表示服务提供商是否支持这种操作
	Supported *bool `json:"supported,omitempty"`

	// 一次可操作的最大个数
	MaxOperations *int32 `json:"maxOperations,omitempty"`

	// 最大有效载荷量
	MaxPayloadSize *int32 `json:"maxPayloadSize,omitempty"`
}

func (o BulkDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BulkDto struct{}"
	}

	return strings.Join([]string{"BulkDto", string(data)}, " ")
}
