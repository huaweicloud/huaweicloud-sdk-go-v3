package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListShareCacheGroupsRecord 共享缓存域名记录
type ListShareCacheGroupsRecord struct {

	// 共享缓存组关联域名
	DomainName *string `json:"domain_name,omitempty"`
}

func (o ListShareCacheGroupsRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListShareCacheGroupsRecord struct{}"
	}

	return strings.Join([]string{"ListShareCacheGroupsRecord", string(data)}, " ")
}
