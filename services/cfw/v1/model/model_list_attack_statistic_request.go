package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAttackStatisticRequest Request Object
type ListAttackStatisticRequest struct {

	// **参数解释**： 防火墙ID，用户创建防火墙实例后产生的唯一ID，配置后可区分不同防火墙，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	FwInstanceId string `json:"fw_instance_id"`

	// **参数解释**： 时间范围  **约束限制**： 不涉及 **取值范围**： 0为近一时 1近一天 2近七天   **默认取值**： 不涉及
	Range *ListAttackStatisticRequestRange `json:"range,omitempty"`

	// **参数解释**： 日志类型 **约束限制**： 不涉及 **取值范围**： internet为南北向日志、nat为nat场景日志，vpc为东西向日志，vgw为vgw场景日志 **默认取值**： 不涉及
	LogType ListAttackStatisticRequestLogType `json:"log_type"`

	// **参数解释**： 会话方向 **约束限制**： 不涉及 **取值范围**： in2out为出云方向 out2in为入云方向 **默认取值**： 不涉及
	Direction *ListAttackStatisticRequestDirection `json:"direction,omitempty"`

	// **参数解释**： 开始时间 **约束限制**： 不涉及 **取值范围**： 毫秒级时间戳 **默认取值**： 不涉及
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**： 结束时间 **约束限制**： 不涉及 **取值范围**： 毫秒级时间戳 **默认取值**： 不涉及
	EndTime *int64 `json:"end_time,omitempty"`

	// **参数解释**： VGW ID **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	VgwId *[]string `json:"vgw_id,omitempty"`

	// **参数解释**： 聚合类型 **约束限制**： 不涉及 **取值范围**： dst TOP攻击目的统计 src TOP攻击来源统计 **默认取值**： 不涉及
	Item ListAttackStatisticRequestItem `json:"item"`
}

func (o ListAttackStatisticRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAttackStatisticRequest struct{}"
	}

	return strings.Join([]string{"ListAttackStatisticRequest", string(data)}, " ")
}

type ListAttackStatisticRequestRange struct {
	value int32
}

type ListAttackStatisticRequestRangeEnum struct {
	E_0 ListAttackStatisticRequestRange
	E_1 ListAttackStatisticRequestRange
	E_2 ListAttackStatisticRequestRange
}

func GetListAttackStatisticRequestRangeEnum() ListAttackStatisticRequestRangeEnum {
	return ListAttackStatisticRequestRangeEnum{
		E_0: ListAttackStatisticRequestRange{
			value: 0,
		}, E_1: ListAttackStatisticRequestRange{
			value: 1,
		}, E_2: ListAttackStatisticRequestRange{
			value: 2,
		},
	}
}

func (c ListAttackStatisticRequestRange) Value() int32 {
	return c.value
}

func (c ListAttackStatisticRequestRange) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAttackStatisticRequestRange) UnmarshalJSON(b []byte) error {
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

type ListAttackStatisticRequestLogType struct {
	value string
}

type ListAttackStatisticRequestLogTypeEnum struct {
	INTERNET ListAttackStatisticRequestLogType
	NAT      ListAttackStatisticRequestLogType
	VPC      ListAttackStatisticRequestLogType
	VGW      ListAttackStatisticRequestLogType
}

func GetListAttackStatisticRequestLogTypeEnum() ListAttackStatisticRequestLogTypeEnum {
	return ListAttackStatisticRequestLogTypeEnum{
		INTERNET: ListAttackStatisticRequestLogType{
			value: "internet",
		},
		NAT: ListAttackStatisticRequestLogType{
			value: "nat",
		},
		VPC: ListAttackStatisticRequestLogType{
			value: "vpc",
		},
		VGW: ListAttackStatisticRequestLogType{
			value: "vgw",
		},
	}
}

func (c ListAttackStatisticRequestLogType) Value() string {
	return c.value
}

func (c ListAttackStatisticRequestLogType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAttackStatisticRequestLogType) UnmarshalJSON(b []byte) error {
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

type ListAttackStatisticRequestDirection struct {
	value string
}

type ListAttackStatisticRequestDirectionEnum struct {
	IN2OUT ListAttackStatisticRequestDirection
	OUT2IN ListAttackStatisticRequestDirection
}

func GetListAttackStatisticRequestDirectionEnum() ListAttackStatisticRequestDirectionEnum {
	return ListAttackStatisticRequestDirectionEnum{
		IN2OUT: ListAttackStatisticRequestDirection{
			value: "in2out",
		},
		OUT2IN: ListAttackStatisticRequestDirection{
			value: "out2in",
		},
	}
}

func (c ListAttackStatisticRequestDirection) Value() string {
	return c.value
}

func (c ListAttackStatisticRequestDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAttackStatisticRequestDirection) UnmarshalJSON(b []byte) error {
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

type ListAttackStatisticRequestItem struct {
	value string
}

type ListAttackStatisticRequestItemEnum struct {
	DST ListAttackStatisticRequestItem
	SRC ListAttackStatisticRequestItem
}

func GetListAttackStatisticRequestItemEnum() ListAttackStatisticRequestItemEnum {
	return ListAttackStatisticRequestItemEnum{
		DST: ListAttackStatisticRequestItem{
			value: "dst",
		},
		SRC: ListAttackStatisticRequestItem{
			value: "src",
		},
	}
}

func (c ListAttackStatisticRequestItem) Value() string {
	return c.value
}

func (c ListAttackStatisticRequestItem) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAttackStatisticRequestItem) UnmarshalJSON(b []byte) error {
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
