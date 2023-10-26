package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReqConfigHpcCacheBackend hpc缓存型配置信息请求
type ReqConfigHpcCacheBackend struct {
	UpdateHpcCache *ReqUpdateHpcCacheInfo `json:"update_hpc_cache"`
}

func (o ReqConfigHpcCacheBackend) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReqConfigHpcCacheBackend struct{}"
	}

	return strings.Join([]string{"ReqConfigHpcCacheBackend", string(data)}, " ")
}
