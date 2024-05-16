package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RecoveryBackupTarget struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o RecoveryBackupTarget) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecoveryBackupTarget struct{}"
	}

	return strings.Join([]string{"RecoveryBackupTarget", string(data)}, " ")
}
