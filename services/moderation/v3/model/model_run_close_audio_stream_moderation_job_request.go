package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunCloseAudioStreamModerationJobRequest Request Object
type RunCloseAudioStreamModerationJobRequest struct {

	// 创建作业成功时，接口返回的job_id。
	JobId string `json:"job_id"`
}

func (o RunCloseAudioStreamModerationJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunCloseAudioStreamModerationJobRequest struct{}"
	}

	return strings.Join([]string{"RunCloseAudioStreamModerationJobRequest", string(data)}, " ")
}
