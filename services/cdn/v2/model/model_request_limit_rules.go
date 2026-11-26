package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RequestLimitRules **参数解释：** 请求限速配置 **约束限制：** 不涉及
type RequestLimitRules struct {

	// **参数解释：** 是否开启请求限速 - 此为旧接口中的参数，后续将下线  **约束限制：** - request_limit_rules字段置空时代表关闭请求限速功能 - 只支持on，off无效  **取值范围：** on **默认取值：** 不涉及
	Status *string `json:"status,omitempty"`

	// **参数解释：** 优先级 > 数值越大优先级越高  **约束限制：** 优先级不能相同 **取值范围：** 1-1000 **默认取值：** 不涉及
	Priority int32 `json:"priority"`

	// **参数解释：** 匹配类型 **约束限制：** 不涉及 **取值范围：** - all：所有文件 - catalog：目录  **默认取值：** 不涉及
	MatchType string `json:"match_type"`

	// **参数解释：** 匹配类型值 **约束限制：** - 当match_type为all时传空值，例如：\"\" - 当match_type为catalog时传目录地址，以“/”作为首字符,例如：\"/test\" - 值为catalog的时候必填  **取值范围：** 不涉及 **默认取值：** 不涉及
	MatchValue *string `json:"match_value,omitempty"`

	// **参数解释：** 限速方式 **约束限制：** 当前仅支持按流量大小限速，取值为size **取值范围：** 不涉及 **默认取值：** 不涉及
	Type string `json:"type"`

	// **参数解释：** 限速条件 > 例如：type=size，limit_rate_after=50表示从传输50个字节后开始限速且限速值为limit_rate_value  **约束限制：** 不涉及 **取值范围：** 0-1073741824，单位：byte **默认取值：** 不涉及
	LimitRateAfter int64 `json:"limit_rate_after"`

	// **参数解释：** 限速值，即达到限速条件后的最大访问速度 **约束限制：** 不涉及 **取值范围：** 0-104857600，单位：Bps **默认取值：** 不涉及
	LimitRateValue int32 `json:"limit_rate_value"`

	// **参数解释：** 指明限速的时段 **约束限制：** - 按照每天24小时设置限速时段，格式为：HHMM-HHMM（HH为时，MM为分，时区为UTC+8） - 多个时段限速时用“,”分隔，例如：0100-0200,2200-2300 - 最多可配置10个时段  **取值范围：** 不涉及 **默认取值：** 0000-2400
	LimitTime *string `json:"limit_time,omitempty"`
}

func (o RequestLimitRules) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RequestLimitRules struct{}"
	}

	return strings.Join([]string{"RequestLimitRules", string(data)}, " ")
}
