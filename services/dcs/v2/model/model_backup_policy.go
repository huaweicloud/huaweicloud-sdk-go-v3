package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BackupPolicy struct {

	// 备份类型。 - auto：自动备份 - manual：手动备份
	BackupType string `json:"backup_type"`

	// 当backup_type设置为auto时，该参数为必填。 保留天数，单位：天，取值范围：1-7。
	SaveDays *int32 `json:"save_days,omitempty"`

	PeriodicalBackupPlan *BackupPlan `json:"periodical_backup_plan,omitempty"`
}

func (o BackupPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupPolicy struct{}"
	}

	return strings.Join([]string{"BackupPolicy", string(data)}, " ")
}
