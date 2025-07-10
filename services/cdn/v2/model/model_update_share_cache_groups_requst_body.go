package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateShareCacheGroupsRequstBody **参数解释：** 共享缓存组配置 **约束限制：** 不涉及
type UpdateShareCacheGroupsRequstBody struct {

	// **参数解释：** 关联域名 **约束限制：** - 必须传入主域名 - 最多可包含50个关联域名
	ShareCacheRecords *[]ShareCacheGroupsRecord `json:"share_cache_records,omitempty"`
}

func (o UpdateShareCacheGroupsRequstBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateShareCacheGroupsRequstBody struct{}"
	}

	return strings.Join([]string{"UpdateShareCacheGroupsRequstBody", string(data)}, " ")
}
