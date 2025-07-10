package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListShareCacheGroupsResponse Response Object
type ListShareCacheGroupsResponse struct {

	// 共享缓存配置总数量
	Total *int32 `json:"total,omitempty"`

	// 共享缓存配置
	ShareCacheGroups *[]ListShareCacheGroupsConfig `json:"share_cache_groups,omitempty"`
	HttpStatusCode   int                           `json:"-"`
}

func (o ListShareCacheGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListShareCacheGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListShareCacheGroupsResponse", string(data)}, " ")
}
