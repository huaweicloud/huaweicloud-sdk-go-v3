package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 作业创建参数
type AudioModerationResultRequestParams struct {

	// 创建作业时传的event_type参数
	EventType *string `json:"event_type,omitempty"`

	Data *AudioModerationResultRequestParamsData `json:"data,omitempty"`

	// 创建作业时传的callback参数
	Callback *string `json:"callback,omitempty"`

	// 创建作业时传的categories参数
	Categories *[]string `json:"categories,omitempty"`
}

func (o AudioModerationResultRequestParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AudioModerationResultRequestParams struct{}"
	}

	return strings.Join([]string{"AudioModerationResultRequestParams", string(data)}, " ")
}
