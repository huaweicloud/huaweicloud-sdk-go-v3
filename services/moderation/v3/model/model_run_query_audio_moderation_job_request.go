package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunQueryAudioModerationJobRequest Request Object
type RunQueryAudioModerationJobRequest struct {

	// 创建作业成功时，接口返回的job_id。
	JobId string `json:"job_id"`
}

func (o RunQueryAudioModerationJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunQueryAudioModerationJobRequest struct{}"
	}

	return strings.Join([]string{"RunQueryAudioModerationJobRequest", string(data)}, " ")
}
