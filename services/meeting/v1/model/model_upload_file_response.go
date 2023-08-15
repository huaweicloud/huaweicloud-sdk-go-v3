package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadFileResponse Response Object
type UploadFileResponse struct {

	// 文件Id。
	FileId *string `json:"fileId,omitempty"`

	ImageModeration *ImageModerationResult `json:"imageModeration,omitempty"`
	HttpStatusCode  int                    `json:"-"`
}

func (o UploadFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadFileResponse struct{}"
	}

	return strings.Join([]string{"UploadFileResponse", string(data)}, " ")
}
