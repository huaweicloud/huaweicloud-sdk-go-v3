package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunCloseVideoStreamModerationJobResponse Response Object
type RunCloseVideoStreamModerationJobResponse struct {

	// 关闭视频流内容审核作业结果
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunCloseVideoStreamModerationJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunCloseVideoStreamModerationJobResponse struct{}"
	}

	return strings.Join([]string{"RunCloseVideoStreamModerationJobResponse", string(data)}, " ")
}
