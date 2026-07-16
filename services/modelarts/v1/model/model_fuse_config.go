package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FuseConfig **参数解释：** 熔断配置。 **取值范围：** 不涉及。
type FuseConfig struct {

	// **参数解释：** 错误率熔断开关。 **约束限制：** 不涉及。 **取值范围：** * true：开启错误率熔断。 * false：不打开错误率熔断。 **默认取值：** false：不打开错误率熔断。
	ErrorRateFuseEnable *bool `json:"error_rate_fuse_enable,omitempty"`

	// **参数解释：** 错误率熔断阈值。 **约束限制：** 不涉及。 **取值范围：** (0, 1]（最多支持2位小数，小数点后第3位做四舍五入处理）。 **默认取值：** 不涉及。
	ErrorRateThreshold *float32 `json:"error_rate_threshold,omitempty"`
}

func (o FuseConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FuseConfig struct{}"
	}

	return strings.Join([]string{"FuseConfig", string(data)}, " ")
}
