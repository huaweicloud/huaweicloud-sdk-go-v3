package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountdownRequestBody **参数解释**： 创建构建任务接口请求体。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type CountdownRequestBody struct {

	// **参数解释**： 服务类型。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ServiceType string `json:"service_type"`

	// **参数解释**： 资源id。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ResourceId string `json:"resource_id"`

	// **参数解释**： 提醒日期。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ReminderDay string `json:"reminder_day"`
}

func (o CountdownRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountdownRequestBody struct{}"
	}

	return strings.Join([]string{"CountdownRequestBody", string(data)}, " ")
}
