package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadSslCertResponse Response Object
type DownloadSslCertResponse struct {

	// SSL证书文件名。
	FileName *string `json:"file_name,omitempty"`

	// SSL证书下载链接。
	Link *string `json:"link,omitempty"`

	// 保存SSL证书的obs桶名。
	BucketName     *string `json:"bucket_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DownloadSslCertResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadSslCertResponse struct{}"
	}

	return strings.Join([]string{"DownloadSslCertResponse", string(data)}, " ")
}
