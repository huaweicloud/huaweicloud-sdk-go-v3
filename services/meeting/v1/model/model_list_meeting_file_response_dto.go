package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IdeaHub、终端“查看会议纪要列表”响应
type ListMeetingFileResponseDto struct {

	// 会议纪要文件码
	FileCode *string `json:"fileCode,omitempty" xml:"fileCode"`

	// 文件主题
	Topic *string `json:"topic,omitempty" xml:"topic"`

	// 文件Id
	FileId *string `json:"fileId,omitempty" xml:"fileId"`

	// 文件名
	FileName *string `json:"fileName,omitempty" xml:"fileName"`

	// 文件大小，单位字节
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize"`

	// 缩略图文件Id
	ThumbnailFileId *string `json:"thumbnailFileId,omitempty" xml:"thumbnailFileId"`

	// 缩略图文件名
	ThumbnailFileName *string `json:"thumbnailFileName,omitempty" xml:"thumbnailFileName"`

	// 缩略图文件大小，单位字节
	ThumbnailFileSize *int64 `json:"thumbnailFileSize,omitempty" xml:"thumbnailFileSize"`

	// pdf文件Id
	PdfFileId *string `json:"pdfFileId,omitempty" xml:"pdfFileId"`

	// pdf文件名
	PdfFileName *string `json:"pdfFileName,omitempty" xml:"pdfFileName"`

	// pdf文件大小，单位字节
	PdfFileSize *int64 `json:"pdfFileSize,omitempty" xml:"pdfFileSize"`

	// 用户ID
	UserId *string `json:"userId,omitempty" xml:"userId"`

	// 文件创建时间戳
	CreationTimestamp *int64 `json:"creationTimestamp,omitempty" xml:"creationTimestamp"`
}

func (o ListMeetingFileResponseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMeetingFileResponseDto struct{}"
	}

	return strings.Join([]string{"ListMeetingFileResponseDto", string(data)}, " ")
}
