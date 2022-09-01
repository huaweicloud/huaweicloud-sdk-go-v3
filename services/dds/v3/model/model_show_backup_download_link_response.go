package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowBackupDownloadLinkResponse struct {

	// 备份文件信息。
	Files *[]GetBackupDownloadLinkResponseBodyFiles `json:"files,omitempty" xml:"files"`

	// OBS桶名。
	Bucket         *string `json:"bucket,omitempty" xml:"bucket"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowBackupDownloadLinkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupDownloadLinkResponse struct{}"
	}

	return strings.Join([]string{"ShowBackupDownloadLinkResponse", string(data)}, " ")
}
