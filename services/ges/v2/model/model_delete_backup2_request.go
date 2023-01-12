package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteBackup2Request struct {

	// 图ID。
	GraphId string `json:"graph_id"`

	// 图备份ID。
	BackupId string `json:"backup_id"`
}

func (o DeleteBackup2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBackup2Request struct{}"
	}

	return strings.Join([]string{"DeleteBackup2Request", string(data)}, " ")
}
