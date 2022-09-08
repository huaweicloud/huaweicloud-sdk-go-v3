package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchStopJobsResponse struct {

	// 批量暂停返回列表
	Results *[]PauseJobResp `json:"results,omitempty"`

	// 总数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchStopJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStopJobsResponse struct{}"
	}

	return strings.Join([]string{"BatchStopJobsResponse", string(data)}, " ")
}
