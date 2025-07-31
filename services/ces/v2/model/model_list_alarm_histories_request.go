package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAlarmHistoriesRequest Request Object
type ListAlarmHistoriesRequest struct {

	// **参数解释**： 告警ID列表。告警ID：以al开头，后跟22位由字母或数字组成的字符串。 **约束限制**： 列表最大长度为50。
	AlarmId *[]string `json:"alarm_id,omitempty"`

	// **参数解释**： 告警记录ID。 **约束限制**： 不涉及。 **取值范围**： 以ah开头，后跟22位由字母或数字组成的字符串，字符串长度为24。 **默认取值**： 不涉及。
	RecordId *string `json:"record_id,omitempty"`

	// **参数解释**： 告警规则名称。 **约束限制**： 不涉及。 **取值范围**： 最大128字符长度。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 告警规则状态列表。告警规则状态：枚举值，ok为正常，alarm为告警，invalid为已失效。 **约束限制**： 列表长度最大为3。
	Status *[]ListAlarmHistoriesRequestStatus `json:"status,omitempty"`

	// **参数解释**： 告警级别。 **约束限制**： 不涉及。 **取值范围**： 枚举值，1为紧急，2为重要，3为次要，4为提示。 **默认取值**： 不涉及。
	Level *int32 `json:"level,omitempty"`

	// **参数解释**： 查询服务的命名空间，各服务命名空间请参考“[服务命名空间](ces_03_0059.xml)”。 **约束限制**： 不涉及。 **取值范围**： 格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_。字符串的长度必须在 3 到 32个字符之间。 **默认取值**： 不涉及。
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**： 告警资源ID。 **约束限制**： 不涉及。 **取值范围**： 多维度情况按字母升序排列并使用逗号分隔。最大字符长度为2048。 **默认取值**： 不涉及。
	ResourceId *string `json:"resource_id,omitempty"`

	// **参数解释**： 查询告警记录的起始更新时间，例如：2022-02-10T10:05:46+08:00。 **约束限制**： 不涉及。 **取值范围**： 最大字符长度为64。 **默认取值**： 不涉及。
	From *string `json:"from,omitempty"`

	// **参数解释**： 查询告警记录的截止更新时间，例如：2022-02-10T10:05:47+08:00。 **约束限制**： 不涉及。 **取值范围**： 最大字符长度为64。 **默认取值**： 不涉及。
	To *string `json:"to,omitempty"`

	// **参数解释**： 告警类型。 **约束限制**： 不涉及。 **取值范围**： 枚举值。event:查询事件类型告警，metric:查询指标类型告警。 **默认取值**： 不涉及。
	AlarmType *ListAlarmHistoriesRequestAlarmType `json:"alarm_type,omitempty"`

	// **参数解释**： 查询告警记录的起始创建时间，例如：2022-02-10T10:05:46+08:00。 **约束限制**： 不涉及。 **取值范围**： 最大字符长度为64。 **默认取值**： 不涉及。
	CreateTimeFrom *string `json:"create_time_from,omitempty"`

	// **参数解释**： 查询告警记录的截止创建时间，例如：2022-02-10T10:05:47+08:00。 **约束限制**： 不涉及。 **取值范围**： 最大字符长度为64。 **默认取值**： 不涉及。
	CreateTimeTo *string `json:"create_time_to,omitempty"`

	// **参数解释**： 分页偏移量。 **约束限制**： 不涉及。 **取值范围**： 整数，最小值为0，最大值为1000000000。 **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 分页偏移量。 **约束限制**： 不涉及。 **取值范围**： 整数，最小值为1，最大值为100。 **默认取值**： 100
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 按关键字排序。 **约束限制**： 不涉及。 **取值范围**： 枚举值。first_alarm_time:告警产生时间, update_time:更新时间, alarm_level:告警级别，record_id表记录主键。 **默认取值**： update_time
	OrderBy *ListAlarmHistoriesRequestOrderBy `json:"order_by,omitempty"`
}

func (o ListAlarmHistoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmHistoriesRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmHistoriesRequest", string(data)}, " ")
}

type ListAlarmHistoriesRequestStatus struct {
	value string
}

type ListAlarmHistoriesRequestStatusEnum struct {
	OK      ListAlarmHistoriesRequestStatus
	ALARM   ListAlarmHistoriesRequestStatus
	INVALID ListAlarmHistoriesRequestStatus
}

func GetListAlarmHistoriesRequestStatusEnum() ListAlarmHistoriesRequestStatusEnum {
	return ListAlarmHistoriesRequestStatusEnum{
		OK: ListAlarmHistoriesRequestStatus{
			value: "ok",
		},
		ALARM: ListAlarmHistoriesRequestStatus{
			value: "alarm",
		},
		INVALID: ListAlarmHistoriesRequestStatus{
			value: "invalid",
		},
	}
}

func (c ListAlarmHistoriesRequestStatus) Value() string {
	return c.value
}

func (c ListAlarmHistoriesRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlarmHistoriesRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListAlarmHistoriesRequestAlarmType struct {
	value string
}

type ListAlarmHistoriesRequestAlarmTypeEnum struct {
	EVENT  ListAlarmHistoriesRequestAlarmType
	METRIC ListAlarmHistoriesRequestAlarmType
}

func GetListAlarmHistoriesRequestAlarmTypeEnum() ListAlarmHistoriesRequestAlarmTypeEnum {
	return ListAlarmHistoriesRequestAlarmTypeEnum{
		EVENT: ListAlarmHistoriesRequestAlarmType{
			value: "event",
		},
		METRIC: ListAlarmHistoriesRequestAlarmType{
			value: "metric",
		},
	}
}

func (c ListAlarmHistoriesRequestAlarmType) Value() string {
	return c.value
}

func (c ListAlarmHistoriesRequestAlarmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlarmHistoriesRequestAlarmType) UnmarshalJSON(b []byte) error {
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

type ListAlarmHistoriesRequestOrderBy struct {
	value string
}

type ListAlarmHistoriesRequestOrderByEnum struct {
	FIRST_ALARM_TIME ListAlarmHistoriesRequestOrderBy
	UPDATE_TIME      ListAlarmHistoriesRequestOrderBy
	ALARM_LEVEL      ListAlarmHistoriesRequestOrderBy
	RECORD_ID        ListAlarmHistoriesRequestOrderBy
}

func GetListAlarmHistoriesRequestOrderByEnum() ListAlarmHistoriesRequestOrderByEnum {
	return ListAlarmHistoriesRequestOrderByEnum{
		FIRST_ALARM_TIME: ListAlarmHistoriesRequestOrderBy{
			value: "first_alarm_time",
		},
		UPDATE_TIME: ListAlarmHistoriesRequestOrderBy{
			value: "update_time",
		},
		ALARM_LEVEL: ListAlarmHistoriesRequestOrderBy{
			value: "alarm_level",
		},
		RECORD_ID: ListAlarmHistoriesRequestOrderBy{
			value: "record_id",
		},
	}
}

func (c ListAlarmHistoriesRequestOrderBy) Value() string {
	return c.value
}

func (c ListAlarmHistoriesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlarmHistoriesRequestOrderBy) UnmarshalJSON(b []byte) error {
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
