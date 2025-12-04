package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StatisticsTimelineItem 不同键值对应的时间线
type StatisticsTimelineItem struct {

	// **参数解释：** 键值标识，用于区分不同的防护统计类型 **约束限制：** 不涉及 **取值范围：**  - ACCESS:请求总量  - CRAWLER:Bot攻击防护  - ATTACK:攻击总量  - WEB_ATTACK:Web基础防护  - PRECISE:精准防护  - CC:CC攻击防护 **默认取值：** 不涉及
	Key *StatisticsTimelineItemKey `json:"key,omitempty"`

	// 对应键值的时间线统计数据
	Timeline *[]TimeLineItem `json:"timeline,omitempty"`
}

func (o StatisticsTimelineItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatisticsTimelineItem struct{}"
	}

	return strings.Join([]string{"StatisticsTimelineItem", string(data)}, " ")
}

type StatisticsTimelineItemKey struct {
	value string
}

type StatisticsTimelineItemKeyEnum struct {
	ACCESS     StatisticsTimelineItemKey
	CRAWLER    StatisticsTimelineItemKey
	ATTACK     StatisticsTimelineItemKey
	WEB_ATTACK StatisticsTimelineItemKey
	PRECISE    StatisticsTimelineItemKey
	CC         StatisticsTimelineItemKey
}

func GetStatisticsTimelineItemKeyEnum() StatisticsTimelineItemKeyEnum {
	return StatisticsTimelineItemKeyEnum{
		ACCESS: StatisticsTimelineItemKey{
			value: "ACCESS",
		},
		CRAWLER: StatisticsTimelineItemKey{
			value: "CRAWLER",
		},
		ATTACK: StatisticsTimelineItemKey{
			value: "ATTACK",
		},
		WEB_ATTACK: StatisticsTimelineItemKey{
			value: "WEB_ATTACK",
		},
		PRECISE: StatisticsTimelineItemKey{
			value: "PRECISE",
		},
		CC: StatisticsTimelineItemKey{
			value: "CC",
		},
	}
}

func (c StatisticsTimelineItemKey) Value() string {
	return c.value
}

func (c StatisticsTimelineItemKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StatisticsTimelineItemKey) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
