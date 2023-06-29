package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMeetingFileResponse Response Object
type ShowMeetingFileResponse struct {

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

	// 文件url。
	FileUrl *string `json:"fileUrl,omitempty"`

	// 缩略图文件url。
	ThumbnailUrl *string `json:"thumbnailUrl,omitempty"`

	// pdf文件url。
	PdfUrl *string `json:"pdfUrl,omitempty"`

	// 文件创建时间戳。
	CreationTimestamp *int64 `json:"creationTimestamp,omitempty"`
	HttpStatusCode    int    `json:"-"`
}

func (o ShowMeetingFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMeetingFileResponse struct{}"
	}

	return strings.Join([]string{"ShowMeetingFileResponse", string(data)}, " ")
}
