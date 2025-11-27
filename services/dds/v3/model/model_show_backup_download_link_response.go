package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBackupDownloadLinkResponse Response Object
type ShowBackupDownloadLinkResponse struct {

	// 备份文件信息。
	Files *[]GetBackupDownloadLinkResponseBodyFiles `json:"files,omitempty"`

	// OBS桶名。
	Bucket *string `json:"bucket,omitempty"`

	// 组ID。
	GroupId *string `json:"group_id,omitempty"`

	// 组名。
	GroupName      *string `json:"group_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowBackupDownloadLinkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupDownloadLinkResponse struct{}"
	}

	return strings.Join([]string{"ShowBackupDownloadLinkResponse", string(data)}, " ")
}
