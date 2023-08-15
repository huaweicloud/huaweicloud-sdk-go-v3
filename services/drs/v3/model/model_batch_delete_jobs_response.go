package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteJobsResponse Response Object
type BatchDeleteJobsResponse struct {

	// 批量结束任务或删除任务的响应体集合。
	Results *[]DeleteJobResp `json:"results,omitempty"`

	// 总数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchDeleteJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteJobsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteJobsResponse", string(data)}, " ")
}
