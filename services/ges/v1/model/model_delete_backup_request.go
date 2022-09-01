package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteBackupRequest struct {

	// 图备份ID。
	BackupId string `json:"backup_id" xml:"backup_id"`

	// 图ID。
	GraphId string `json:"graph_id" xml:"graph_id"`
}

func (o DeleteBackupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBackupRequest struct{}"
	}

	return strings.Join([]string{"DeleteBackupRequest", string(data)}, " ")
}
