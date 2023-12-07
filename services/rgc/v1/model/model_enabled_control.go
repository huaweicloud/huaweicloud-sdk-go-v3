package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnabledControl 开启控制策略信息。
type EnabledControl struct {

	// 管理账号ID。
	ManageAccountId *string `json:"manage_account_id,omitempty"`

	// 控制策略标识。
	ControlIdentifier *string `json:"control_identifier,omitempty"`

	// 控制策略名称。
	Name *string `json:"name,omitempty"`

	// 控制策略描述。
	Description *string `json:"description,omitempty"`

	// 控制策略目标。
	ControlObjective *string `json:"control_objective,omitempty"`

	// 控制策略类型。包括主动性控制策略Proactive、检测性控制策略Detective、预防性控制策略Preventive。
	Behavior *string `json:"behavior,omitempty"`

	// 账号的创建来源，包括CUSTOM和RGC。
	Owner *string `json:"owner,omitempty"`

	// regional/global。
	RegionalPreference *string `json:"regional_preference,omitempty"`
}

func (o EnabledControl) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnabledControl struct{}"
	}

	return strings.Join([]string{"EnabledControl", string(data)}, " ")
}
