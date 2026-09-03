package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyTrainingQuotaItem **参数解释**：训练作业配额项。
type ModifyTrainingQuotaItem struct {

	// **参数解释**：配额的资源类型。 **约束限制**：不涉及。 **取值范围**：枚举值如下： - job-num: 作业个数配额 - visual-job-num: 可视化作业个数配额 - job-retention-enabled: 用户级作业自动老化开关 - job-num-quota-notify: 配额告警SMN通知配置 **默认取值**：不涉及。
	Resource string `json:"resource"`

	// **参数解释**：配额个数。 **约束限制**：取值约束因资源类型而异：job-retention-enabled取值0（关闭）或1（开启）；job-num-quota-notify固定为0，通知主题URN存于extra_info；其余资源类型要求不小于1。 **取值范围**：0 ~ 2147483647。 **默认取值**：不涉及。
	Quota int32 `json:"quota"`

	// **参数解释**：已使用的个数。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Used *int32 `json:"used,omitempty"`

	// **参数解释**：配额的额外信息。 **约束限制**：当resource为job-num-quota-notify时，该字段存储SMN通知主题URN。 **取值范围**：不涉及。 **默认取值**：不涉及。
	ExtraInfo *string `json:"extra_info,omitempty"`
}

func (o ModifyTrainingQuotaItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyTrainingQuotaItem struct{}"
	}

	return strings.Join([]string{"ModifyTrainingQuotaItem", string(data)}, " ")
}
