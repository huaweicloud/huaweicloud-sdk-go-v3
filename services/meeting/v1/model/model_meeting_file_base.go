package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MeetingFileBase 会议纪要文件信息。
type MeetingFileBase struct {

	// 会议纪要文件码。
	FileCode *string `json:"fileCode,omitempty"`

	// 文件主题。
	Topic *string `json:"topic,omitempty"`

	// 文件Id。
	FileId *string `json:"fileId,omitempty"`

	// 文件名。
	FileName *string `json:"fileName,omitempty"`

	// 文件大小，单位字节。
	FileSize *int64 `json:"fileSize,omitempty"`

	// 缩略图文件Id。
	ThumbnailFileId *string `json:"thumbnailFileId,omitempty"`

	// 缩略图文件名。
	ThumbnailFileName *string `json:"thumbnailFileName,omitempty"`

	// 缩略图文件大小，单位字节。
	ThumbnailFileSize *int64 `json:"thumbnailFileSize,omitempty"`

	// pdf文件Id。
	PdfFileId *string `json:"pdfFileId,omitempty"`

	// pdf文件名。
	PdfFileName *string `json:"pdfFileName,omitempty"`

	// pdf文件大小，单位字节。
	PdfFileSize *int64 `json:"pdfFileSize,omitempty"`
}

func (o MeetingFileBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MeetingFileBase struct{}"
	}

	return strings.Join([]string{"MeetingFileBase", string(data)}, " ")
}
