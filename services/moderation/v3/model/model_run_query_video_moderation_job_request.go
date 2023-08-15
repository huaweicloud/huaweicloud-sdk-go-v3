package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunQueryVideoModerationJobRequest Request Object
type RunQueryVideoModerationJobRequest struct {

	// 创建作业成功时，接口返回的job_id。
	JobId string `json:"job_id"`
}

func (o RunQueryVideoModerationJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunQueryVideoModerationJobRequest struct{}"
	}

	return strings.Join([]string{"RunQueryVideoModerationJobRequest", string(data)}, " ")
}
