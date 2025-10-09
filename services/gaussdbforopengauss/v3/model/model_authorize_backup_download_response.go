package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuthorizeBackupDownloadResponse Response Object
type AuthorizeBackupDownloadResponse struct {

	// OBS桶名。
	Bucket *string `json:"bucket,omitempty"`

	// 通过OBS Browser+下载备份文件的路径名称。
	FilePaths      *[]string `json:"file_paths,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o AuthorizeBackupDownloadResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizeBackupDownloadResponse struct{}"
	}

	return strings.Join([]string{"AuthorizeBackupDownloadResponse", string(data)}, " ")
}
