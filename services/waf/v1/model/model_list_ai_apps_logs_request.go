package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAiAppsLogsRequest Request Object
type ListAiAppsLogsRequest struct {

	// **参数解释：** APP ID筛选 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	AppId *string `json:"app_id,omitempty"`

	// **参数解释：** 请求ID精确筛选，返回匹配该request_id的检测日志 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	RequestId *string `json:"request_id,omitempty"`

	// **参数解释：** 检测类型筛选，可选值：text, image, text_image **约束限制：** 不涉及 **取值范围：** - text: 文本检测 - image: 图片检测 - text_image: 文本图片混合检测 **默认取值：** 不涉及
	DetectType *string `json:"detect_type,omitempty"`

	// **参数解释：** 攻击类型筛选，多个用逗号分隔 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	AttackTypes *[]string `json:"attack_types,omitempty"`

	// **参数解释：** 执行动作筛选，可选值：pass, block, log, desensitize **约束限制：** 不涉及 **取值范围：** - pass: 放行 - block: 拦截 - log: 记录 - desensitize: 脱敏 **默认取值：** 不涉及
	Suggestion *string `json:"suggestion,omitempty"`

	// **参数解释：** APP名称筛选，按应用名称精确匹配查询 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	AppName *string `json:"app_name,omitempty"`

	// **参数解释：** 开始时间，时间格式为毫秒时间戳 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	From *int64 `json:"from,omitempty"`

	// **参数解释：** 截止时间，时间格式为毫秒时间戳 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	To *int64 `json:"to,omitempty"`

	// **参数解释：** 每页数量，默认10，最大1000 **约束限制：** 最大1000 **取值范围：** 1-1000 **默认取值：** 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 偏移量，默认0 **约束限制：** 不涉及 **取值范围：** >=0 **默认取值：** 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 游标标记，标记到ES的request_id字段，传入时忽略offset参数 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Marker *string `json:"marker,omitempty"`

	// **参数解释：** 排序字段，仅支持timestamp **约束限制：** 仅支持timestamp **取值范围：** - timestamp: 按时间排序 **默认取值：** timestamp
	SortBy *string `json:"sort_by,omitempty"`

	// **参数解释：** 排序方向，可选值：asc, desc **约束限制：** 不涉及 **取值范围：** - asc: 升序 - desc: 降序 **默认取值：** desc
	SortDirection *string `json:"sort_direction,omitempty"`
}

func (o ListAiAppsLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAiAppsLogsRequest struct{}"
	}

	return strings.Join([]string{"ListAiAppsLogsRequest", string(data)}, " ")
}
