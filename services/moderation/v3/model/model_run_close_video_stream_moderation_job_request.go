package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunCloseVideoStreamModerationJobRequest Request Object
type RunCloseVideoStreamModerationJobRequest struct {

	// 创建作业成功时，接口返回的job_id。
	JobId string `json:"job_id"`
}

func (o RunCloseVideoStreamModerationJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunCloseVideoStreamModerationJobRequest struct{}"
	}

	return strings.Join([]string{"RunCloseVideoStreamModerationJobRequest", string(data)}, " ")
}
