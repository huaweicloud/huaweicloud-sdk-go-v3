package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLargeFileResponse Response Object
type CreateLargeFileResponse struct {

	// 文件ID。
	FileId *string `json:"file_id,omitempty"`

	// 文件分段上传URL。
	UploadUrls *[]string `json:"upload_urls,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateLargeFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLargeFileResponse struct{}"
	}

	return strings.Join([]string{"CreateLargeFileResponse", string(data)}, " ")
}
