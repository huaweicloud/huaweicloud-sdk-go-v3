package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ObjectTaskResult struct {

	// 工作流中任务类型
	TaskType *string `json:"task_type,omitempty"`

	TranscodeTask *TranscodeTask `json:"transcode_task,omitempty"`

	ThumbnailTask *ThumbnailTask `json:"thumbnail_task,omitempty"`

	ImageSpriteTask *ImageSpriteTask `json:"image_sprite_task,omitempty"`
}

func (o ObjectTaskResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObjectTaskResult struct{}"
	}

	return strings.Join([]string{"ObjectTaskResult", string(data)}, " ")
}
