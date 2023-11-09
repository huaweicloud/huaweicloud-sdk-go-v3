package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DocumentQueryResponseRequestParams 作业创建参数
type DocumentQueryResponseRequestParams struct {
	Data *DocumentQueryResponseRequestParamsData `json:"data"`

	// 创建作业时传的event_type参数
	EventType string `json:"event_type"`

	// 创建作业时传的image_categories参数
	ImageCategories *[]string `json:"image_categories,omitempty"`

	// 创建作业时传的text_categories参数
	TextCategories *[]string `json:"text_categories,omitempty"`

	// 创建作业时传的video_image_categories参数
	VideoImageCategories *[]string `json:"video_image_categories,omitempty"`

	// 创建作业时传的audio_categories参数
	AudioCategories *[]string `json:"audio_categories,omitempty"`

	// 创建作业时传的callback参数
	Callback *string `json:"callback,omitempty"`
}

func (o DocumentQueryResponseRequestParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DocumentQueryResponseRequestParams struct{}"
	}

	return strings.Join([]string{"DocumentQueryResponseRequestParams", string(data)}, " ")
}
