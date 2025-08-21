package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TargetControl 治理策略概要。
type TargetControl struct {

	// 管理员账号ID。
	ManageAccountId *string `json:"manage_account_id,omitempty"`

	// 控制策略标识。
	ControlIdentifier *string `json:"control_identifier,omitempty"`

	// 控制策略启用状态。
	State *string `json:"state,omitempty"`

	// 控制策略当前版本号。
	Version *string `json:"version,omitempty"`

	// 控制策略名称。
	Name *string `json:"name,omitempty"`

	// 控制策略描述信息。
	Description *string `json:"description,omitempty"`

	// 控制策略目标。
	ControlObjective *string `json:"control_objective,omitempty"`

	// 控制策略类型。包括主动性控制策略Proactive、检测性控制策略Detective、预防性控制策略Preventive。
	Behavior *string `json:"behavior,omitempty"`

	// 控制策略来源。
	Owner *string `json:"owner,omitempty"`

	// 区域选项，取值有两种分别是：区域的regional和全局的global。
	RegionalPreference *string `json:"regional_preference,omitempty"`

	// 控制策略必须性。
	Guidance *string `json:"guidance,omitempty"`

	// 控制策略所属服务。
	Service *string `json:"service,omitempty"`

	// 策略类别。
	Implementation *string `json:"implementation,omitempty"`
}

func (o TargetControl) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TargetControl struct{}"
	}

	return strings.Join([]string{"TargetControl", string(data)}, " ")
}
