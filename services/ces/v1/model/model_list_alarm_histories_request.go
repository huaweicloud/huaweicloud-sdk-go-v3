package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAlarmHistoriesRequest Request Object
type ListAlarmHistoriesRequest struct {

	// **参数解释**： 当前资源所在分组信息 **约束限制**： 不涉及。 **取值范围**： 以rg开头，后跟22位由字母或数字组成的字符串，字符长度为24 **默认取值**： 不涉及
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释**： 告警规则ID **约束限制**： 不涉及 **取值范围**： 以al开头，后跟22位由字母或数字组成的字符串，字符长度为24 **默认取值**： 不涉及
	AlarmId *string `json:"alarm_id,omitempty"`

	// **参数解释**： 告警规则名称 **约束限制**： 不涉及 **取值范围**： 只能包含0-9/a-z/A-Z/_/-或汉字，长度[1, 128]个字符 **默认取值**： 不涉及
	AlarmName *string `json:"alarm_name,omitempty"`

	// **参数解释**： 告警状态。 **约束限制**： 不涉及 **取值范围**： 枚举值： - ok：正常 - alarm：告警 - insufficient_data：数据不足 - invalid：已失效 **默认取值**： 不涉及
	AlarmStatus *ListAlarmHistoriesRequestAlarmStatus `json:"alarm_status,omitempty"`

	// **参数解释**： 告警历史的告警级别，值为1,2,3,4 **约束限制**： 不涉及 **取值范围**： 枚举值： - 1：紧急 - 2：重要 - 3：次要 - 4：提示 **默认取值**： 不涉及
	AlarmLevel *ListAlarmHistoriesRequestAlarmLevel `json:"alarm_level,omitempty"`

	// **参数解释**： 查询服务的命名空间，各服务命名空间请参阅[[支持监控的服务列表](https://support.huaweicloud.com/api-ces/ces_03_0059.html)](tag:hc)[[支持监控的服务列表](https://support.huaweicloud.com/intl/en-us/api-ces/ces_03_0059.html)](tag:hk)[[支持监控的服务列表](https://support.huaweicloud.com/eu/en-us/api-ces/ces_03_0059.html)](tag:hws_eu)[[支持监控的服务列表](ces_03_0059.xml)](tag:ax,cmcc,ctc,dt,dt_test,hcso_dt,fcs,fcs_vm,mix,g42,hk_g42,hk_sbc,hk_tm,hk_vdf,hws_ocb,ocb,sbc,srg)。 **约束限制**： 不涉及 **取值范围**： 格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_。字符串的长度必须在 3 到 32个字符之间 **默认取值**： 不涉及
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**： 通过时间筛选traces的起始时间(包括传入时间)，为timestamp **约束限制**： 不涉及 **取值范围**： 长度为[1,13]个字符 **默认取值**： 不涉及
	From *string `json:"from,omitempty"`

	// **参数解释**： 通过时间筛选traces的终止时间(包括传入时间)，为timestamp **约束限制**： 不涉及 **取值范围**： 长度为[1,13]个字符 **默认取值**： 不涉及
	To *string `json:"to,omitempty"`

	// **参数解释**： 分页查询时查询的起始位置，表示从第几条数据开始 **约束限制**： 不涉及。 **取值范围**： 大于等于0的整数 **默认取值**： 0
	Start *string `json:"start,omitempty"`

	// **参数解释**： 本次查询的最大条目数 **约束限制**： 不涉及。 **取值范围**： 取值范围[1,100] **默认取值**： 100
	Limit *string `json:"limit,omitempty"`
}

func (o ListAlarmHistoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmHistoriesRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmHistoriesRequest", string(data)}, " ")
}

type ListAlarmHistoriesRequestAlarmStatus struct {
	value string
}

type ListAlarmHistoriesRequestAlarmStatusEnum struct {
	OK                ListAlarmHistoriesRequestAlarmStatus
	ALARM             ListAlarmHistoriesRequestAlarmStatus
	INSUFFICIENT_DATA ListAlarmHistoriesRequestAlarmStatus
	INVALID           ListAlarmHistoriesRequestAlarmStatus
}

func GetListAlarmHistoriesRequestAlarmStatusEnum() ListAlarmHistoriesRequestAlarmStatusEnum {
	return ListAlarmHistoriesRequestAlarmStatusEnum{
		OK: ListAlarmHistoriesRequestAlarmStatus{
			value: "ok",
		},
		ALARM: ListAlarmHistoriesRequestAlarmStatus{
			value: "alarm",
		},
		INSUFFICIENT_DATA: ListAlarmHistoriesRequestAlarmStatus{
			value: "insufficient_data",
		},
		INVALID: ListAlarmHistoriesRequestAlarmStatus{
			value: "invalid",
		},
	}
}

func (c ListAlarmHistoriesRequestAlarmStatus) Value() string {
	return c.value
}

func (c ListAlarmHistoriesRequestAlarmStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlarmHistoriesRequestAlarmStatus) UnmarshalJSON(b []byte) error {
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

type ListAlarmHistoriesRequestAlarmLevel struct {
	value int32
}

type ListAlarmHistoriesRequestAlarmLevelEnum struct {
	E_1 ListAlarmHistoriesRequestAlarmLevel
	E_2 ListAlarmHistoriesRequestAlarmLevel
	E_3 ListAlarmHistoriesRequestAlarmLevel
	E_4 ListAlarmHistoriesRequestAlarmLevel
}

func GetListAlarmHistoriesRequestAlarmLevelEnum() ListAlarmHistoriesRequestAlarmLevelEnum {
	return ListAlarmHistoriesRequestAlarmLevelEnum{
		E_1: ListAlarmHistoriesRequestAlarmLevel{
			value: 1,
		}, E_2: ListAlarmHistoriesRequestAlarmLevel{
			value: 2,
		}, E_3: ListAlarmHistoriesRequestAlarmLevel{
			value: 3,
		}, E_4: ListAlarmHistoriesRequestAlarmLevel{
			value: 4,
		},
	}
}

func (c ListAlarmHistoriesRequestAlarmLevel) Value() int32 {
	return c.value
}

func (c ListAlarmHistoriesRequestAlarmLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlarmHistoriesRequestAlarmLevel) UnmarshalJSON(b []byte) error {
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
