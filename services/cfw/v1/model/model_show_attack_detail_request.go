package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAttackDetailRequest Request Object
type ShowAttackDetailRequest struct {

	// **参数解释**： 防火墙ID，用户创建防火墙实例后产生的唯一ID，配置后可区分不同防火墙，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	FwInstanceId string `json:"fw_instance_id"`

	// **参数解释**： 时间范围  **约束限制**： 不涉及 **取值范围**： 0为近一时 1近一天 2近七天   **默认取值**： 不涉及
	Range *ShowAttackDetailRequestRange `json:"range,omitempty"`

	// **参数解释**： 日志类型 **约束限制**： 不涉及 **取值范围**： internet为南北向日志、nat为nat场景日志，vpc为东西向日志，vgw为vgw场景日志 **默认取值**： 不涉及
	LogType ShowAttackDetailRequestLogType `json:"log_type"`

	// **参数解释**： 开始时间 **约束限制**： 不涉及 **取值范围**： 毫秒级时间戳 **默认取值**： 不涉及
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**： 结束时间 **约束限制**： 不涉及 **取值范围**： 毫秒级时间戳 **默认取值**： 不涉及
	EndTime *int64 `json:"end_time,omitempty"`

	// **参数解释**： VGW ID **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	VgwId *[]string `json:"vgw_id,omitempty"`

	// **参数解释**： 动作 **约束限制**： 不涉及 **取值范围**： 0 全部 1 拦截 **默认取值**： 不涉及
	Action ShowAttackDetailRequestAction `json:"action"`

	// **参数解释**： 聚合类型 **约束限制**： 不涉及 **取值范围**： src_region_id TOP外部攻击来源地区 attack_type 攻击类型 in_src_ip TOP内部攻击来源IP out_src_ip TOP外部攻击来源IP dst_port TOP被攻击端口 dst_ip TOP攻击目的IP attack_rule TOP攻击规则 src_ip TOP攻击来源IP **默认取值**： 不涉及
	Item ShowAttackDetailRequestItem `json:"item"`

	// **参数解释**： 统计对象 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Value string `json:"value"`
}

func (o ShowAttackDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAttackDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowAttackDetailRequest", string(data)}, " ")
}

type ShowAttackDetailRequestRange struct {
	value int32
}

type ShowAttackDetailRequestRangeEnum struct {
	E_0 ShowAttackDetailRequestRange
	E_1 ShowAttackDetailRequestRange
	E_2 ShowAttackDetailRequestRange
}

func GetShowAttackDetailRequestRangeEnum() ShowAttackDetailRequestRangeEnum {
	return ShowAttackDetailRequestRangeEnum{
		E_0: ShowAttackDetailRequestRange{
			value: 0,
		}, E_1: ShowAttackDetailRequestRange{
			value: 1,
		}, E_2: ShowAttackDetailRequestRange{
			value: 2,
		},
	}
}

func (c ShowAttackDetailRequestRange) Value() int32 {
	return c.value
}

func (c ShowAttackDetailRequestRange) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAttackDetailRequestRange) UnmarshalJSON(b []byte) error {
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

type ShowAttackDetailRequestLogType struct {
	value string
}

type ShowAttackDetailRequestLogTypeEnum struct {
	INTERNET ShowAttackDetailRequestLogType
	NAT      ShowAttackDetailRequestLogType
	VPC      ShowAttackDetailRequestLogType
	VGW      ShowAttackDetailRequestLogType
}

func GetShowAttackDetailRequestLogTypeEnum() ShowAttackDetailRequestLogTypeEnum {
	return ShowAttackDetailRequestLogTypeEnum{
		INTERNET: ShowAttackDetailRequestLogType{
			value: "internet",
		},
		NAT: ShowAttackDetailRequestLogType{
			value: "nat",
		},
		VPC: ShowAttackDetailRequestLogType{
			value: "vpc",
		},
		VGW: ShowAttackDetailRequestLogType{
			value: "vgw",
		},
	}
}

func (c ShowAttackDetailRequestLogType) Value() string {
	return c.value
}

func (c ShowAttackDetailRequestLogType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAttackDetailRequestLogType) UnmarshalJSON(b []byte) error {
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

type ShowAttackDetailRequestAction struct {
	value int32
}

type ShowAttackDetailRequestActionEnum struct {
	E_0 ShowAttackDetailRequestAction
	E_1 ShowAttackDetailRequestAction
}

func GetShowAttackDetailRequestActionEnum() ShowAttackDetailRequestActionEnum {
	return ShowAttackDetailRequestActionEnum{
		E_0: ShowAttackDetailRequestAction{
			value: 0,
		}, E_1: ShowAttackDetailRequestAction{
			value: 1,
		},
	}
}

func (c ShowAttackDetailRequestAction) Value() int32 {
	return c.value
}

func (c ShowAttackDetailRequestAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAttackDetailRequestAction) UnmarshalJSON(b []byte) error {
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

type ShowAttackDetailRequestItem struct {
	value string
}

type ShowAttackDetailRequestItemEnum struct {
	SRC_REGION_ID ShowAttackDetailRequestItem
	ATTACK_TYPE   ShowAttackDetailRequestItem
	IN_SRC_IP     ShowAttackDetailRequestItem
	OUT_SRC_IP    ShowAttackDetailRequestItem
	DST_PORT      ShowAttackDetailRequestItem
	DST_IP        ShowAttackDetailRequestItem
	ATTACK_RULE   ShowAttackDetailRequestItem
	SRC_IP        ShowAttackDetailRequestItem
}

func GetShowAttackDetailRequestItemEnum() ShowAttackDetailRequestItemEnum {
	return ShowAttackDetailRequestItemEnum{
		SRC_REGION_ID: ShowAttackDetailRequestItem{
			value: "src_region_id",
		},
		ATTACK_TYPE: ShowAttackDetailRequestItem{
			value: "attack_type",
		},
		IN_SRC_IP: ShowAttackDetailRequestItem{
			value: "in_src_ip",
		},
		OUT_SRC_IP: ShowAttackDetailRequestItem{
			value: "out_src_ip",
		},
		DST_PORT: ShowAttackDetailRequestItem{
			value: "dst_port",
		},
		DST_IP: ShowAttackDetailRequestItem{
			value: "dst_ip",
		},
		ATTACK_RULE: ShowAttackDetailRequestItem{
			value: "attack_rule",
		},
		SRC_IP: ShowAttackDetailRequestItem{
			value: "src_ip",
		},
	}
}

func (c ShowAttackDetailRequestItem) Value() string {
	return c.value
}

func (c ShowAttackDetailRequestItem) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAttackDetailRequestItem) UnmarshalJSON(b []byte) error {
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
