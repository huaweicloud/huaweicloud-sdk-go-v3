package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 提交异步任务响应体。
type AsyncCommitJobResp struct {

	// 批量异步任务ID。
	AsyncJobId string `json:"async_job_id"`

	// 批量异步任务状态。
	Status string `json:"status"`

	// 提交指定ID批量异步任务结果信息。
	Msg string `json:"msg"`
}

func (o AsyncCommitJobResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AsyncCommitJobResp struct{}"
	}

	return strings.Join([]string{"AsyncCommitJobResp", string(data)}, " ")
}
