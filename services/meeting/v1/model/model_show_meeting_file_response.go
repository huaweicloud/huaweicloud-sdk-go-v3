package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowMeetingFileResponse struct {

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

	// 文件url
	FileUrl *string `json:"fileUrl,omitempty" xml:"fileUrl"`

	// 缩略图文件url
	ThumbnailUrl *string `json:"thumbnailUrl,omitempty" xml:"thumbnailUrl"`

	// pdf文件url
	PdfUrl *string `json:"pdfUrl,omitempty" xml:"pdfUrl"`

	// 文件创建时间戳
	CreationTimestamp *int64 `json:"creationTimestamp,omitempty" xml:"creationTimestamp"`
	HttpStatusCode    int    `json:"-"`
}

func (o ShowMeetingFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMeetingFileResponse struct{}"
	}

	return strings.Join([]string{"ShowMeetingFileResponse", string(data)}, " ")
}
