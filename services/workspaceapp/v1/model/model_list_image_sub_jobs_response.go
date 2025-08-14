package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageSubJobsResponse Response Object
type ListImageSubJobsResponse struct {

	// 总数。
	Count *int32 `json:"count,omitempty"`

	// 任务明细信息列表，返回列表条目数量上限为分页的最大上限值。
	Items          *[]ImageJobDetailInfo `json:"items,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListImageSubJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageSubJobsResponse struct{}"
	}

	return strings.Join([]string{"ListImageSubJobsResponse", string(data)}, " ")
}
