package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAttackTopRequest Request Object
type ShowAttackTopRequest struct {

	// **参数解释**： 防火墙ID，用户创建防火墙实例后产生的唯一ID，配置后可区分不同防火墙，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	FwInstanceId string `json:"fw_instance_id"`

	// **参数解释**： 时间范围  **约束限制**： 不涉及 **取值范围**： 0为近一时 1近一天 2近七天   **默认取值**： 不涉及
	Range *ShowAttackTopRequestRange `json:"range,omitempty"`

	// **参数解释**： 日志类型 **约束限制**： 不涉及 **取值范围**： internet为南北向日志、nat为nat场景日志，vpc为东西向日志，vgw为vgw场景日志 **默认取值**： 不涉及
	LogType ShowAttackTopRequestLogType `json:"log_type"`

	// **参数解释**： 开始时间 **约束限制**： 不涉及 **取值范围**： 毫秒级时间戳 **默认取值**： 不涉及
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**： 结束时间 **约束限制**： 不涉及 **取值范围**： 毫秒级时间戳 **默认取值**： 不涉及
	EndTime *int64 `json:"end_time,omitempty"`

	// **参数解释**： VGW ID **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	VgwId *[]string `json:"vgw_id,omitempty"`

	// **参数解释**： 动作 **约束限制**： 不涉及 **取值范围**： 0 全部 1 拦截 **默认取值**： 不涉及
	Action ShowAttackTopRequestAction `json:"action"`

	// **参数解释**： 聚合类型 **约束限制**： 不涉及 **取值范围**： src_region_id TOP外部攻击来源地区 attack_type 攻击类型 in_src_ip TOP内部攻击来源IP out_src_ip TOP外部攻击来源IP dst_port TOP被攻击端口 dst_ip TOP攻击目的IP attack_rule TOP攻击规则 src_ip TOP攻击来源IP level TOP威胁等级 **默认取值**： 不涉及
	Item []string `json:"item"`

	// **参数解释**： 聚合条数 **约束限制**： 不涉及 **取值范围**： 0到10条 **默认取值**： 5
	Size *int32 `json:"size,omitempty"`
}

func (o ShowAttackTopRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAttackTopRequest struct{}"
	}

	return strings.Join([]string{"ShowAttackTopRequest", string(data)}, " ")
}

type ShowAttackTopRequestRange struct {
	value int32
}

type ShowAttackTopRequestRangeEnum struct {
	E_0 ShowAttackTopRequestRange
	E_1 ShowAttackTopRequestRange
	E_2 ShowAttackTopRequestRange
}

func GetShowAttackTopRequestRangeEnum() ShowAttackTopRequestRangeEnum {
	return ShowAttackTopRequestRangeEnum{
		E_0: ShowAttackTopRequestRange{
			value: 0,
		}, E_1: ShowAttackTopRequestRange{
			value: 1,
		}, E_2: ShowAttackTopRequestRange{
			value: 2,
		},
	}
}

func (c ShowAttackTopRequestRange) Value() int32 {
	return c.value
}

func (c ShowAttackTopRequestRange) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAttackTopRequestRange) UnmarshalJSON(b []byte) error {
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

type ShowAttackTopRequestLogType struct {
	value string
}

type ShowAttackTopRequestLogTypeEnum struct {
	INTERNET ShowAttackTopRequestLogType
	NAT      ShowAttackTopRequestLogType
	VPC      ShowAttackTopRequestLogType
	VGW      ShowAttackTopRequestLogType
}

func GetShowAttackTopRequestLogTypeEnum() ShowAttackTopRequestLogTypeEnum {
	return ShowAttackTopRequestLogTypeEnum{
		INTERNET: ShowAttackTopRequestLogType{
			value: "internet",
		},
		NAT: ShowAttackTopRequestLogType{
			value: "nat",
		},
		VPC: ShowAttackTopRequestLogType{
			value: "vpc",
		},
		VGW: ShowAttackTopRequestLogType{
			value: "vgw",
		},
	}
}

func (c ShowAttackTopRequestLogType) Value() string {
	return c.value
}

func (c ShowAttackTopRequestLogType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAttackTopRequestLogType) UnmarshalJSON(b []byte) error {
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

type ShowAttackTopRequestAction struct {
	value int32
}

type ShowAttackTopRequestActionEnum struct {
	E_0 ShowAttackTopRequestAction
	E_1 ShowAttackTopRequestAction
}

func GetShowAttackTopRequestActionEnum() ShowAttackTopRequestActionEnum {
	return ShowAttackTopRequestActionEnum{
		E_0: ShowAttackTopRequestAction{
			value: 0,
		}, E_1: ShowAttackTopRequestAction{
			value: 1,
		},
	}
}

func (c ShowAttackTopRequestAction) Value() int32 {
	return c.value
}

func (c ShowAttackTopRequestAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAttackTopRequestAction) UnmarshalJSON(b []byte) error {
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
