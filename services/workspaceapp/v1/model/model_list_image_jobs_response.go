package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageJobsResponse Response Object
type ListImageJobsResponse struct {

	// 总数。
	Count *int32 `json:"count,omitempty"`

	// 任务信息列表，返回列表条目数量上限为分页的最大上限值。
	Items          *[]ImageJobInfo `json:"items,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListImageJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageJobsResponse struct{}"
	}

	return strings.Join([]string{"ListImageJobsResponse", string(data)}, " ")
}
