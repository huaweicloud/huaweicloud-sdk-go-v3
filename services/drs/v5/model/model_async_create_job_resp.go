package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AsyncCreateJobResp 创建异步任务响应体。
type AsyncCreateJobResp struct {

	// 批量异步创建的任务ID。
	AsyncJobId string `json:"async_job_id"`
}

func (o AsyncCreateJobResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AsyncCreateJobResp struct{}"
	}

	return strings.Join([]string{"AsyncCreateJobResp", string(data)}, " ")
}
