package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportBackup2Request Request Object
type ImportBackup2Request struct {

	// 图ID。
	GraphId string `json:"graph_id"`

	Body *ImportBackupReq `json:"body,omitempty"`
}

func (o ImportBackup2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportBackup2Request struct{}"
	}

	return strings.Join([]string{"ImportBackup2Request", string(data)}, " ")
}
