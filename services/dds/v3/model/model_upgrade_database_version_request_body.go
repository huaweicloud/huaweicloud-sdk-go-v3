package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpgradeDatabaseVersionRequestBody struct {

	// **参数解释：** 升级模式。 **约束限制：** 不涉及。 **取值范围：** minimized_interrupt_time：表示中断时间最短优先模式：升级过程对业务影响相对较小的升级方式。 minimized_upgrade_time：表示升级时长最短优先模式：升级过程时长相对较快的升级方式。 **默认取值：** minimized_interrupt_time。
	UpgradeMode *string `json:"upgrade_mode,omitempty"`

	// **参数解释：** 实例是否在可维护时间窗内自动升级。 **约束限制：** 不涉及。 **取值范围：** true：表示延迟升级，实例将在可维护时间窗内自动升级。 false：表示立即升级。 **默认取值：** false。
	IsDelayed *bool `json:"is_delayed,omitempty"`
}

func (o UpgradeDatabaseVersionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeDatabaseVersionRequestBody struct{}"
	}

	return strings.Join([]string{"UpgradeDatabaseVersionRequestBody", string(data)}, " ")
}
