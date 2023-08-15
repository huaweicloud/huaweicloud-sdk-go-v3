package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRacksRequest Request Object
type ListRacksRequest struct {

	// 每页的数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页标识
	Marker *string `json:"marker,omitempty"`

	// 排序字段
	SortKey *[]string `json:"sort_key,omitempty"`

	// 排序方向，取值范围： - desc：降序 - acs：升序
	SortDir *[]string `json:"sort_dir,omitempty"`

	// 边缘小站ID
	EdgeSiteId *string `json:"edge_site_id,omitempty"`
}

func (o ListRacksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRacksRequest struct{}"
	}

	return strings.Join([]string{"ListRacksRequest", string(data)}, " ")
}
