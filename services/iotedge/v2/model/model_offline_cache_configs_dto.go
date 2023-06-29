package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OfflineCacheConfigsDto 节点离线缓存配置
type OfflineCacheConfigsDto struct {

	// 数据上报优先级，可选项：realtime_first实时数据优先 sequential按时序上报，默认realtime_first
	PublishOrder *string `json:"publish_order,omitempty"`

	// 节点离线缓存数据的储存天数，默认7，取值范围-1~14，-1表示存储天数没有限制
	Period *int32 `json:"period,omitempty"`

	// 节点离线缓存容量，单位MB，默认2048，取值范围500-8192
	Capacity *int32 `json:"capacity,omitempty"`
}

func (o OfflineCacheConfigsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OfflineCacheConfigsDto struct{}"
	}

	return strings.Join([]string{"OfflineCacheConfigsDto", string(data)}, " ")
}
