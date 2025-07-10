package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListShareCacheGroupsConfig 共享缓存配置
type ListShareCacheGroupsConfig struct {

	// 共享缓存ID
	Id *string `json:"id,omitempty"`

	// 共享缓存组名
	GroupName *string `json:"group_name,omitempty"`

	// 共享缓存创建时间
	CreateTime *int32 `json:"create_time,omitempty"`

	// 共享缓存域名记录
	ShareCacheRecords *[]ListShareCacheGroupsRecord `json:"share_cache_records,omitempty"`

	// 共享缓存组主域名
	PrimaryDomain *string `json:"primary_domain,omitempty"`
}

func (o ListShareCacheGroupsConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListShareCacheGroupsConfig struct{}"
	}

	return strings.Join([]string{"ListShareCacheGroupsConfig", string(data)}, " ")
}
