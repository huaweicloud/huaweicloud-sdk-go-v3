package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TargetConfig 实例规格变更时，实例的目标规格配置。
type TargetConfig struct {
	Flavor *Flavor `json:"flavor,omitempty"`

	// **参数说明**：实例的付费方式。 **取值范围**： - prePaid：包年/包月 - postPaid：按需计费
	ChargeMode *string `json:"charge_mode,omitempty"`
}

func (o TargetConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TargetConfig struct{}"
	}

	return strings.Join([]string{"TargetConfig", string(data)}, " ")
}
