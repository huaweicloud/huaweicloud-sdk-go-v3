package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReqUpdateHpcCacheInfo 后端挂载信息
type ReqUpdateHpcCacheInfo struct {

	// 配置hpc缓存型的动作，如initialize_overlay
	Action string `json:"action"`

	Data *ReqUpdateHpcCacheData `json:"data"`
}

func (o ReqUpdateHpcCacheInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReqUpdateHpcCacheInfo struct{}"
	}

	return strings.Join([]string{"ReqUpdateHpcCacheInfo", string(data)}, " ")
}
