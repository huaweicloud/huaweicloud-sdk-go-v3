package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 文件上传信息。
type OpenUploadFileInfo struct {

	// 文件Id。
	FileId *string `json:"fileId,omitempty"`

	ImageModeration *ImageModerationResult `json:"imageModeration,omitempty"`
}

func (o OpenUploadFileInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenUploadFileInfo struct{}"
	}

	return strings.Join([]string{"OpenUploadFileInfo", string(data)}, " ")
}
