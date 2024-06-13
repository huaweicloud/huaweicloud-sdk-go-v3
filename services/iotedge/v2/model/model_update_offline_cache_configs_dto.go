package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateOfflineCacheConfigsDto 节点离线缓存配置
type UpdateOfflineCacheConfigsDto struct {

	// 节点离线缓存容量，单位MB，默认2048，取值范围500-65536
	Capacity *int32 `json:"capacity,omitempty"`
}

func (o UpdateOfflineCacheConfigsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOfflineCacheConfigsDto struct{}"
	}

	return strings.Join([]string{"UpdateOfflineCacheConfigsDto", string(data)}, " ")
}
