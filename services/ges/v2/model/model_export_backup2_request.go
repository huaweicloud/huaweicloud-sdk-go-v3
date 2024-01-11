package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportBackup2Request Request Object
type ExportBackup2Request struct {

	// 图ID。
	GraphId string `json:"graph_id"`

	Body *ExportBackupReq `json:"body,omitempty"`
}

func (o ExportBackup2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportBackup2Request struct{}"
	}

	return strings.Join([]string{"ExportBackup2Request", string(data)}, " ")
}
