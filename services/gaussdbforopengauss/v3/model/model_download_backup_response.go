package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadBackupResponse Response Object
type DownloadBackupResponse struct {

	// 备份文件信息。
	Files *[]DownloadObject `json:"files,omitempty"`

	// 桶名。
	Bucket         *string `json:"bucket,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DownloadBackupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadBackupResponse struct{}"
	}

	return strings.Join([]string{"DownloadBackupResponse", string(data)}, " ")
}
