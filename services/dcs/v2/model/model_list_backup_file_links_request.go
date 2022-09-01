package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListBackupFileLinksRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 备份记录ID。
	BackupId string `json:"backup_id" xml:"backup_id"`

	Body *DownloadBackupFilesReq `json:"body,omitempty" xml:"body"`
}

func (o ListBackupFileLinksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackupFileLinksRequest struct{}"
	}

	return strings.Join([]string{"ListBackupFileLinksRequest", string(data)}, " ")
}
