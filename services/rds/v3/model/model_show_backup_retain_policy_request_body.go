package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBackupRetainPolicyRequestBody 查询备份保留策略的请求体
type ShowBackupRetainPolicyRequestBody struct {

	// **参数解释**：  实例ID列表，实例ID是实例的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  实例ID只能由英文字母、数字组成，长度为36个字符。  **默认取值**：  不涉及。
	Instanceids *[]string `json:"instanceids,omitempty"`

	// **参数解释**  索引位置，偏移量。  **约束限制**  从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。  **取值范围**  大于等于0的整数。  **默认取值**  0
	Offset int32 `json:"offset"`

	// **参数解释**  查询记录数。  **约束限制**  不能为负数。  **取值范围**  最小值为1，最大值为100。  **默认取值**  10
	Limit int32 `json:"limit"`

	// **参数解释**：  实例状态  **约束限制**：  不涉及。  **取值范围**：  normal、deleted  **默认取值**：  不涉及。
	InstanceStatus *string `json:"instance_status,omitempty"`

	// **参数解释**  查询开始时间。时间指实例的删除时间。  **约束限制**  “begin_time”有值时，“end_time”必选。 “begin_time”有值时，查询实例状态为已删除的实例。  **取值范围**  格式为“yyyy-mm-ddThh:mm:ss±HH:mm”。  其中，T指某个时间的开始；±HH:mm指时区偏移量，例如北京时间偏移显示为+08:00。  **默认取值**  不涉及。
	BeginTime *string `json:"begin_time,omitempty"`

	// **参数解释**  查询结束时间。时间指实例的删除时间  **约束限制**  “end_time”有值时，“begin_time”必选。 “end_time”有值时，查询实例状态为已删除的实例。  **取值范围**  格式为“yyyy-mm-ddThh:mm:ss±HH:mm”，且大于查询开始时间。  其中，T指某个时间的开始；±HH:mm指时区偏移量，例如北京时间偏移显示为+08:00。  **默认取值**  不涉及。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ShowBackupRetainPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupRetainPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"ShowBackupRetainPolicyRequestBody", string(data)}, " ")
}
