package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMissingIndexDetailsResponse Response Object
type ListMissingIndexDetailsResponse struct {

	// 索引缺失明细列表
	DetailList *[]interface{} `json:"detail_list,omitempty"`

	// 总数
	Total *int64 `json:"total,omitempty"`

	// 采集时间
	CollectTime    *int64 `json:"collect_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListMissingIndexDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMissingIndexDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListMissingIndexDetailsResponse", string(data)}, " ")
}
