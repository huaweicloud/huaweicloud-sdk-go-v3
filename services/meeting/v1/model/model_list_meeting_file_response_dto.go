package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IdeaHub、终端“查看会议纪要列表”响应
type ListMeetingFileResponseDto struct {
	// 会议纪要文件码

	FileCode *string `json:"fileCode,omitempty"`
	// 文件主题

	Topic *string `json:"topic,omitempty"`
	// 文件Id

	FileId *string `json:"fileId,omitempty"`
	// 文件名

	FileName *string `json:"fileName,omitempty"`
	// 文件大小，单位字节

	FileSize *int64 `json:"fileSize,omitempty"`
	// 缩略图文件Id

	ThumbnailFileId *string `json:"thumbnailFileId,omitempty"`
	// 缩略图文件名

	ThumbnailFileName *string `json:"thumbnailFileName,omitempty"`
	// 缩略图文件大小，单位字节

	ThumbnailFileSize *int64 `json:"thumbnailFileSize,omitempty"`
	// pdf文件Id

	PdfFileId *string `json:"pdfFileId,omitempty"`
	// pdf文件名

	PdfFileName *string `json:"pdfFileName,omitempty"`
	// pdf文件大小，单位字节

	PdfFileSize *int64 `json:"pdfFileSize,omitempty"`
	// 用户ID

	UserId *string `json:"userId,omitempty"`
	// 文件创建时间戳

	CreationTimestamp *int64 `json:"creationTimestamp,omitempty"`
}

func (o ListMeetingFileResponseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMeetingFileResponseDto struct{}"
	}

	return strings.Join([]string{"ListMeetingFileResponseDto", string(data)}, " ")
}
