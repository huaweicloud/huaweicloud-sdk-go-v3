package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchStartJobsResponse Response Object
type BatchStartJobsResponse struct {

	// 批量启动实时灾备任务返回列表。
	Results *[]StartJobResp `json:"results,omitempty"`

	// 总数。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchStartJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartJobsResponse struct{}"
	}

	return strings.Join([]string{"BatchStartJobsResponse", string(data)}, " ")
}
