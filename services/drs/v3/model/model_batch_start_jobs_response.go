package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchStartJobsResponse struct {

	// 批量启动实时灾备任务返回列表。
	Results *[]StartJobResp `json:"results,omitempty" xml:"results"`

	// 总数。
	Count          *int32 `json:"count,omitempty" xml:"count"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchStartJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartJobsResponse struct{}"
	}

	return strings.Join([]string{"BatchStartJobsResponse", string(data)}, " ")
}
