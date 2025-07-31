package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAttackTotalRequest Request Object
type ShowAttackTotalRequest struct {

	// **参数解释**： 防火墙ID，用户创建防火墙实例后产生的唯一ID，配置后可区分不同防火墙，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	FwInstanceId string `json:"fw_instance_id"`

	// **参数解释**： 时间范围  **约束限制**： 不涉及 **取值范围**： 0为近一时 1近一天 2近七天   **默认取值**： 不涉及
	Range *ShowAttackTotalRequestRange `json:"range,omitempty"`

	// **参数解释**： 日志类型 **约束限制**： 不涉及 **取值范围**： internet为南北向日志、nat为nat场景日志，vpc为东西向日志，vgw为vgw场景日志 **默认取值**： 不涉及
	LogType ShowAttackTotalRequestLogType `json:"log_type"`

	// **参数解释**： 开始时间 **约束限制**： 不涉及 **取值范围**： 毫秒级时间戳 **默认取值**： 不涉及
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**： 结束时间 **约束限制**： 不涉及 **取值范围**： 毫秒级时间戳 **默认取值**： 不涉及
	EndTime *int64 `json:"end_time,omitempty"`

	// **参数解释**： VGW ID **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	VgwId *[]string `json:"vgw_id,omitempty"`
}

func (o ShowAttackTotalRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAttackTotalRequest struct{}"
	}

	return strings.Join([]string{"ShowAttackTotalRequest", string(data)}, " ")
}

type ShowAttackTotalRequestRange struct {
	value int32
}

type ShowAttackTotalRequestRangeEnum struct {
	E_0 ShowAttackTotalRequestRange
	E_1 ShowAttackTotalRequestRange
	E_2 ShowAttackTotalRequestRange
}

func GetShowAttackTotalRequestRangeEnum() ShowAttackTotalRequestRangeEnum {
	return ShowAttackTotalRequestRangeEnum{
		E_0: ShowAttackTotalRequestRange{
			value: 0,
		}, E_1: ShowAttackTotalRequestRange{
			value: 1,
		}, E_2: ShowAttackTotalRequestRange{
			value: 2,
		},
	}
}

func (c ShowAttackTotalRequestRange) Value() int32 {
	return c.value
}

func (c ShowAttackTotalRequestRange) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAttackTotalRequestRange) UnmarshalJSON(b []byte) error {
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

type ShowAttackTotalRequestLogType struct {
	value string
}

type ShowAttackTotalRequestLogTypeEnum struct {
	INTERNET ShowAttackTotalRequestLogType
	NAT      ShowAttackTotalRequestLogType
	VPC      ShowAttackTotalRequestLogType
	VGW      ShowAttackTotalRequestLogType
}

func GetShowAttackTotalRequestLogTypeEnum() ShowAttackTotalRequestLogTypeEnum {
	return ShowAttackTotalRequestLogTypeEnum{
		INTERNET: ShowAttackTotalRequestLogType{
			value: "internet",
		},
		NAT: ShowAttackTotalRequestLogType{
			value: "nat",
		},
		VPC: ShowAttackTotalRequestLogType{
			value: "vpc",
		},
		VGW: ShowAttackTotalRequestLogType{
			value: "vgw",
		},
	}
}

func (c ShowAttackTotalRequestLogType) Value() string {
	return c.value
}

func (c ShowAttackTotalRequestLogType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAttackTotalRequestLogType) UnmarshalJSON(b []byte) error {
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
