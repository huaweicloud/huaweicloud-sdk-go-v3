package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TransferBackupRequestBody 手动转储任务请求体
type TransferBackupRequestBody struct {

	// 备份id列表
	Backups []string `json:"backups"`

	// 实例id
	InstanceId string `json:"instance_id"`

	// 目标OBS桶名
	Destination string `json:"destination"`

	// 目标前缀
	Prefix *string `json:"prefix,omitempty"`
}

func (o TransferBackupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransferBackupRequestBody struct{}"
	}

	return strings.Join([]string{"TransferBackupRequestBody", string(data)}, " ")
}
