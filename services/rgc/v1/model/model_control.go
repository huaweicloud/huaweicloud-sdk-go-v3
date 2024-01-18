package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Control 控制策略。
type Control struct {

	// 控制策略ID。
	Identifier *string `json:"identifier,omitempty"`

	// 控制策略名称。
	Name *string `json:"name,omitempty"`

	// 控制策略描述。
	Description *string `json:"description,omitempty"`

	// 控制策略必须性。
	Guidance *string `json:"guidance,omitempty"`

	// 治理资源。
	Resource *[]string `json:"resource,omitempty"`

	// 治理策略来自的框架。
	Framework *[]string `json:"framework,omitempty"`

	// 控制策略所属服务。
	Service *string `json:"service,omitempty"`

	// 服务控制策略（SCP），配置规则。
	Implementation *string `json:"implementation,omitempty"`

	// 控制策略类型。包括主动性控制策略Proactive、检测性控制策略Detective、预防性控制策略Preventive。
	Behavior *string `json:"behavior,omitempty"`

	// 纳管账号的创建来源，包括CUSTOM和RGC。
	Owner *string `json:"owner,omitempty"`

	// 控制策略严重性。
	Severity *string `json:"severity,omitempty"`

	// 控制策略目标。
	ControlObjective *string `json:"control_objective,omitempty"`

	// 控制策略版本。
	Version *string `json:"version,omitempty"`

	// 控制策略发布时间。
	ReleaseDate *sdktime.SdkTime `json:"release_date,omitempty"`
}

func (o Control) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Control struct{}"
	}

	return strings.Join([]string{"Control", string(data)}, " ")
}
