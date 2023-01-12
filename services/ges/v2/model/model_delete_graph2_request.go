package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteGraph2Request struct {

	// 图ID。
	GraphId string `json:"graph_id"`

	// 删除图后是否保留备份，默认保留1个自动备份和2个手动备份。该查询参数为空时，表示不保留。
	KeepBackup *bool `json:"keep_backup,omitempty"`
}

func (o DeleteGraph2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGraph2Request struct{}"
	}

	return strings.Join([]string{"DeleteGraph2Request", string(data)}, " ")
}
