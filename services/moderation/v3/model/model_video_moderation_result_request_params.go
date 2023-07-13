package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VideoModerationResultRequestParams 作业创建参数
type VideoModerationResultRequestParams struct {
	Data *VideoModerationResultRequestParamsData `json:"data"`

	// 创建作业时传的event_type参数
	EventType string `json:"event_type"`

	// 创建作业时传的image_categories参数
	ImageCategories []string `json:"image_categories"`

	// 创建作业时传的audio_categories参数
	AudioCategories *[]string `json:"audio_categories,omitempty"`

	// 创建作业时传的callback参数
	Callback *string `json:"callback,omitempty"`
}

func (o VideoModerationResultRequestParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoModerationResultRequestParams struct{}"
	}

	return strings.Join([]string{"VideoModerationResultRequestParams", string(data)}, " ")
}
