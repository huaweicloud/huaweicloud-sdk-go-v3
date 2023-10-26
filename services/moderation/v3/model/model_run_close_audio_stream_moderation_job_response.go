package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunCloseAudioStreamModerationJobResponse Response Object
type RunCloseAudioStreamModerationJobResponse struct {

	// 关闭音频流审核作业响应结果
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunCloseAudioStreamModerationJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunCloseAudioStreamModerationJobResponse struct{}"
	}

	return strings.Join([]string{"RunCloseAudioStreamModerationJobResponse", string(data)}, " ")
}
