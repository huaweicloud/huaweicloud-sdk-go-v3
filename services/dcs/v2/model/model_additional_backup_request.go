package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AdditionalBackupRequest struct {

	// 是否高级备份
	IsAdditionalBackup *bool `json:"is_additional_backup,omitempty"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`
}

func (o AdditionalBackupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdditionalBackupRequest struct{}"
	}

	return strings.Join([]string{"AdditionalBackupRequest", string(data)}, " ")
}
