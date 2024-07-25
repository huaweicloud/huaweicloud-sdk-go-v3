package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBackupDownloadLinkResponse Response Object
type ShowBackupDownloadLinkResponse struct {

	// 文件所在的桶名。
	Bucket *string `json:"bucket,omitempty"`

	// 备份包含的文件列表。
	Files          *[]BackupDownloadLink `json:"files,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowBackupDownloadLinkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupDownloadLinkResponse struct{}"
	}

	return strings.Join([]string{"ShowBackupDownloadLinkResponse", string(data)}, " ")
}
