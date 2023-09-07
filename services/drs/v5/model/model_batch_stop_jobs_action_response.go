package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchStopJobsActionResponse Response Object
type BatchStopJobsActionResponse struct {

	// 批量操作任务响应体。
	Jobs           *[]ActionBaseResp `json:"jobs,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o BatchStopJobsActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStopJobsActionResponse struct{}"
	}

	return strings.Join([]string{"BatchStopJobsActionResponse", string(data)}, " ")
}
