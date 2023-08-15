package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Video
type Video struct {

	// 创建时间
	CreateTime string `json:"create_time"`

	// 更新时间
	UpdateTime string `json:"update_time"`

	// 错误信息 如: {\\\"error_code\\\":\\\"0001\\\",\\\"error_msg\\\":\\\"播报内容超过10分钟，请重新调整播报内容。\\\"}
	ErrorMsg *string `json:"error_msg,omitempty"`

	//
	Id string `json:"id"`

	// 视频名称
	Name string `json:"name"`

	// 视频生成进度 0~100
	Progress *int32 `json:"progress,omitempty"`

	// 0：未初始化 1：生成中 2：生成成功 3：生成失败
	Status int32 `json:"status"`

	// 字幕地址
	SubtitleUrl string `json:"subtitle_url"`

	// 视频的obs地址，当视频生成成功时返回
	VideoUrl *string `json:"video_url,omitempty"`

	// 视频截图地址，jpg格式 分辨率480 * 270 当status=2：生成成功时返回
	VideoShot *string `json:"video_shot,omitempty"`
}

func (o Video) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Video struct{}"
	}

	return strings.Join([]string{"Video", string(data)}, " ")
}
