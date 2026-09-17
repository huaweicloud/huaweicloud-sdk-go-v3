package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIndexUsageDetailsResponse Response Object
type ListIndexUsageDetailsResponse struct {

	// 索引缺失明细列表
	DetailList *[]IndexUsageDetail `json:"detail_list,omitempty"`

	// 总数
	Total *int64 `json:"total,omitempty"`

	// 采集时间
	CollectTime    *int64 `json:"collect_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListIndexUsageDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIndexUsageDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListIndexUsageDetailsResponse", string(data)}, " ")
}
