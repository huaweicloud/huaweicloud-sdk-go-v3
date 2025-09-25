package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAlarmHistoryRecordRequest Request Object
type ShowAlarmHistoryRecordRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ShowAlarmHistoryRecordRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 查询开始时间。 **约束限制**: 不涉及。 **取值范围**: 格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。最多可以获取最近7天的数据。 **默认取值**: 不涉及。
	StartTime string `json:"start_time"`

	// **参数解释**: 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询。例如：该参数指定为1，limit指定为10，则只展示第2-11条数据。 **约束限制**: 不涉及。 **取值范围**: [0, 2^31-1] **默认取值**: 默认为0（偏移0条数据，表示从第一条数据开始查询）。
	Offset int32 `json:"offset"`

	// **参数解释**: 查询记录数。例如该参数设定为10，则查询结果最多只显示10条记录。 **约束限制**: 不涉及。 **取值范围**: [1, 50] **默认取值**: 50
	Limit int32 `json:"limit"`

	// **参数解释**: 实例告警等级。 **约束限制**: 不涉及。 **取值范围**: - 1：紧急。 - 2：重要。 - 3：次要。 - 4：提示。  **默认取值**: 1
	Level *ShowAlarmHistoryRecordRequestLevel `json:"level,omitempty"`
}

func (o ShowAlarmHistoryRecordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlarmHistoryRecordRequest struct{}"
	}

	return strings.Join([]string{"ShowAlarmHistoryRecordRequest", string(data)}, " ")
}

type ShowAlarmHistoryRecordRequestXLanguage struct {
	value string
}

type ShowAlarmHistoryRecordRequestXLanguageEnum struct {
	ZH_CN ShowAlarmHistoryRecordRequestXLanguage
	EN_US ShowAlarmHistoryRecordRequestXLanguage
}

func GetShowAlarmHistoryRecordRequestXLanguageEnum() ShowAlarmHistoryRecordRequestXLanguageEnum {
	return ShowAlarmHistoryRecordRequestXLanguageEnum{
		ZH_CN: ShowAlarmHistoryRecordRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowAlarmHistoryRecordRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowAlarmHistoryRecordRequestXLanguage) Value() string {
	return c.value
}

func (c ShowAlarmHistoryRecordRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAlarmHistoryRecordRequestXLanguage) UnmarshalJSON(b []byte) error {
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

type ShowAlarmHistoryRecordRequestLevel struct {
	value int32
}

type ShowAlarmHistoryRecordRequestLevelEnum struct {
	E_1 ShowAlarmHistoryRecordRequestLevel
	E_2 ShowAlarmHistoryRecordRequestLevel
	E_3 ShowAlarmHistoryRecordRequestLevel
	E_4 ShowAlarmHistoryRecordRequestLevel
}

func GetShowAlarmHistoryRecordRequestLevelEnum() ShowAlarmHistoryRecordRequestLevelEnum {
	return ShowAlarmHistoryRecordRequestLevelEnum{
		E_1: ShowAlarmHistoryRecordRequestLevel{
			value: 1,
		}, E_2: ShowAlarmHistoryRecordRequestLevel{
			value: 2,
		}, E_3: ShowAlarmHistoryRecordRequestLevel{
			value: 3,
		}, E_4: ShowAlarmHistoryRecordRequestLevel{
			value: 4,
		},
	}
}

func (c ShowAlarmHistoryRecordRequestLevel) Value() int32 {
	return c.value
}

func (c ShowAlarmHistoryRecordRequestLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAlarmHistoryRecordRequestLevel) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
