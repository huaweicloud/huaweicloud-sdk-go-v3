package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAlarmStatisticsRequest Request Object
type ShowAlarmStatisticsRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ShowAlarmStatisticsRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 查询开始时间。 **约束限制**: 最多可以统计最近7天的数据。 **取值范围**: 格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。 **默认取值**: 不涉及。
	StartTime string `json:"start_time"`

	// **参数解释**: 告警数量最多的实例的个数。 **约束限制**: 不涉及。 **取值范围**: 取值必须大于0。 **默认取值**: 5
	TopNum int32 `json:"top_num"`
}

func (o ShowAlarmStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlarmStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ShowAlarmStatisticsRequest", string(data)}, " ")
}

type ShowAlarmStatisticsRequestXLanguage struct {
	value string
}

type ShowAlarmStatisticsRequestXLanguageEnum struct {
	ZH_CN ShowAlarmStatisticsRequestXLanguage
	EN_US ShowAlarmStatisticsRequestXLanguage
}

func GetShowAlarmStatisticsRequestXLanguageEnum() ShowAlarmStatisticsRequestXLanguageEnum {
	return ShowAlarmStatisticsRequestXLanguageEnum{
		ZH_CN: ShowAlarmStatisticsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowAlarmStatisticsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowAlarmStatisticsRequestXLanguage) Value() string {
	return c.value
}

func (c ShowAlarmStatisticsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAlarmStatisticsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
