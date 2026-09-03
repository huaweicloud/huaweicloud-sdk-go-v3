package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateUploadUrlsResponse Response Object
type CreateUploadUrlsResponse struct {

	// 上传地址列表。
	UploadUrls *[]UploadUrlItem `json:"upload_urls,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateUploadUrlsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUploadUrlsResponse struct{}"
	}

	return strings.Join([]string{"CreateUploadUrlsResponse", string(data)}, " ")
}
