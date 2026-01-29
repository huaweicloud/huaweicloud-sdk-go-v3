package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCollectConfigResponseBodyTarget 目标
type ListCollectConfigResponseBodyTarget struct {

	// 管道
	Pipe *string `json:"pipe,omitempty"`

	// 分片
	Shards float32 `json:"shards,omitempty"`

	// 存储模式
	StorageMode *string `json:"storage_mode,omitempty"`

	// ttl时间
	Ttl float32 `json:"ttl,omitempty"`
}

func (o ListCollectConfigResponseBodyTarget) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCollectConfigResponseBodyTarget struct{}"
	}

	return strings.Join([]string{"ListCollectConfigResponseBodyTarget", string(data)}, " ")
}
