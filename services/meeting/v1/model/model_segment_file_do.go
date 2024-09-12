package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SegmentFileDo 录制切片文件
type SegmentFileDo struct {

	// 会议录制类型，取值范围见数据结构RecordType：AUDIO 纯音频录制，SPEAKER_VIDEO 演讲者视图，SHARE_VIDEO共享屏幕，SPEAKER_SHARE_VIDEO 含演讲者视图的共享屏幕
	RecordType *string `json:"recordType,omitempty"`

	// 录制文件开始时间
	BeginTime *int32 `json:"beginTime,omitempty"`

	// 录制文件结束时间
	EndTime *int32 `json:"endTime,omitempty"`

	// 录制文件时长(秒)
	Duration *int32 `json:"duration,omitempty"`

	// 文件大小(字节数)
	FileSize *int32 `json:"fileSize,omitempty"`

	// 文件hash校验码(SHA256)，64个字符
	Sha256 *string `json:"sha256,omitempty"`

	// 录制文件播放地址，有效期1小时
	PlayUrl *string `json:"playUrl,omitempty"`

	// 录制文件下载地址，有效期1小时
	DownloadUrl *string `json:"downloadUrl,omitempty"`
}

func (o SegmentFileDo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SegmentFileDo struct{}"
	}

	return strings.Join([]string{"SegmentFileDo", string(data)}, " ")
}
