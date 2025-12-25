package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfigurationHistoryRsp struct {

	// **参数解释：** 参数名称。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ParameterName string `json:"parameter_name"`

	// **参数解释：** 参数旧值。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	OldValue string `json:"old_value"`

	// **参数解释：** 参数新值。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	NewValue string `json:"new_value"`

	// **参数解释：** 更新结果。 **约束限制：** 不涉及。 **取值范围：** - SUCCESS：成功。 - FAILED：失败。 **默认取值：** 不涉及。
	UpdateResult string `json:"update_result"`

	// **参数解释：** 是否生效。 **约束限制：** 不涉及。 **取值范围：** - true:已生效。 - false:未生效。 **默认取值：** 不涉及。
	Applied bool `json:"applied"`

	// **参数解释：** 更新时间。 **约束限制：** 不涉及。 **取值范围：** 格式为\"yyyy-MM-ddTHH:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。 **默认取值：** 不涉及。
	UpdatedAt string `json:"updated_at"`

	// **参数解释：** 生效时间。 **约束限制：** 不涉及。 **取值范围：** 格式为\"yyyy-MM-ddTHH:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。 **默认取值：** 不涉及。
	AppliedAt string `json:"applied_at"`
}

func (o ConfigurationHistoryRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationHistoryRsp struct{}"
	}

	return strings.Join([]string{"ConfigurationHistoryRsp", string(data)}, " ")
}
