package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AdditionalObjectProcessReq struct {
	Output *ObsInfo `json:"output"`

	// 转码任务列表
	TranscodeTask *[]ObjectTranscodeTask `json:"transcode_task,omitempty"`

	// 水印设置
	Watermarks *[]WatermarkRequest `json:"watermarks,omitempty"`

	// 转码任务列表
	ThumbnailTask *[]ObjectThumbnailTask `json:"thumbnail_task,omitempty"`

	// 转码任务列表
	ImageSpriteTask *[]ObjectImageSpriteTask `json:"image_sprite_task,omitempty"`

	VideoProcess *VideoProcess `json:"video_process,omitempty"`

	// 回调url
	CallbackUrl *string `json:"callback_url,omitempty"`

	// 自定义上下文，回调时会回调给用户，透传信息
	SessionContext *string `json:"session_context,omitempty"`
}

func (o AdditionalObjectProcessReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdditionalObjectProcessReq struct{}"
	}

	return strings.Join([]string{"AdditionalObjectProcessReq", string(data)}, " ")
}
