package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BackupStrategyDetail 备份策略
type BackupStrategyDetail struct {

	// 策略ID。
	PolicyId *string `json:"policy_id,omitempty"`

	// 策略名称。
	PolicyName *string `json:"policy_name,omitempty"`

	// 执行策略。
	BackupStrategy *string `json:"backup_strategy,omitempty"`

	// 备份类型。
	BackupType *string `json:"backup_type,omitempty"`

	// 备份级别。
	BackupLevel *string `json:"backup_level,omitempty"`

	// 下次触发时间。
	NextFireTime *string `json:"next_fire_time,omitempty"`

	// 更新时间。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o BackupStrategyDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupStrategyDetail struct{}"
	}

	return strings.Join([]string{"BackupStrategyDetail", string(data)}, " ")
}
