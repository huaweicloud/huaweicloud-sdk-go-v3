package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AuditUpgradeStatus struct {

	// 当前版本
	CurrentVersion *string `json:"current_version,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 状态  - 0：正在升级  - 1：满足条件未升级  - 2：不满足升级条件
	Status *int32 `json:"status,omitempty"`

	// 升级步骤
	Steps *[]AuditUpgradeStep `json:"steps,omitempty"`

	// 总量
	Total *int32 `json:"total,omitempty"`

	// 升级版本
	UpgradeVersion *string `json:"upgrade_version,omitempty"`
}

func (o AuditUpgradeStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuditUpgradeStatus struct{}"
	}

	return strings.Join([]string{"AuditUpgradeStatus", string(data)}, " ")
}
