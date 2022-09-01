package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestorePoint struct {

	// 源实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 用于恢复的备份ID。
	BackupId string `json:"backup_id" xml:"backup_id"`
}

func (o RestorePoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestorePoint struct{}"
	}

	return strings.Join([]string{"RestorePoint", string(data)}, " ")
}
