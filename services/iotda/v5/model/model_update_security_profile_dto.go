package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecurityProfileDto 安全态势感知配置
type UpdateSecurityProfileDto struct {

	// 安全态势感知告警级别，CRITICAL：严重告警，MAJOR：重要告警，MINOR：一般告警
	AlarmLevel *string `json:"alarm_level,omitempty"`

	// 安全态势感知项所属安全风险级别；BASIC_SECURITY：基础安全，ADVANCE_SECURITY：高级安全，ULTIMATE_SECURITY：极致安全
	SecurityLevel *string `json:"security_level,omitempty"`

	// 安全态势感知项是否开启
	Enable *bool `json:"enable,omitempty"`

	// 安全态势感知项配置结构体，用于设备侧检测项下发给设备
	Profile *[]SecurityProfile `json:"profile,omitempty"`

	ProfileTargets *SecurityTarget `json:"profile_targets,omitempty"`
}

func (o UpdateSecurityProfileDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityProfileDto struct{}"
	}

	return strings.Join([]string{"UpdateSecurityProfileDto", string(data)}, " ")
}
