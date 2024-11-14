package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDbCacheMappingsResponse Response Object
type ListDbCacheMappingsResponse struct {

	// 总记录数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 内存加速映射信息。
	DbcacheMappings *[]QueryDbCacheMappingResponse `json:"dbcache_mappings,omitempty"`
	HttpStatusCode  int                            `json:"-"`
}

func (o ListDbCacheMappingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDbCacheMappingsResponse struct{}"
	}

	return strings.Join([]string{"ListDbCacheMappingsResponse", string(data)}, " ")
}
