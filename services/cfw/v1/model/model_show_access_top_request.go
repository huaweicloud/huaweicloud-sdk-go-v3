package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAccessTopRequest Request Object
type ShowAccessTopRequest struct {

	// **参数解释**： 防火墙ID，用户创建防火墙实例后产生的唯一ID，配置后可区分不同防火墙，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	FwInstanceId string `json:"fw_instance_id"`

	// **参数解释**： 时间范围  **约束限制**： 不涉及 **取值范围**： 0为近一时 1近一天 2近七天   **默认取值**： 不涉及
	Range *ShowAccessTopRequestRange `json:"range,omitempty"`

	// **参数解释**： 会话方向 **约束限制**： 不涉及 **取值范围**： in2out为出云方向 out2in为入云方向 **默认取值**： 不涉及
	Direction *ShowAccessTopRequestDirection `json:"direction,omitempty"`

	// **参数解释**： 开始时间 **约束限制**： 不涉及 **取值范围**： 毫秒级时间戳 **默认取值**： 不涉及
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**： 结束时间 **约束限制**： 不涉及 **取值范围**： 毫秒级时间戳 **默认取值**： 不涉及
	EndTime *int64 `json:"end_time,omitempty"`

	// **参数解释**： VGW ID **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	VgwId *[]string `json:"vgw_id,omitempty"`

	// **参数解释**： 日志类型 **约束限制**： 不涉及 **取值范围**： internet为南北向日志、nat为nat场景日志，vpc为东西向日志，vgw为vgw场景日志 **默认取值**： 不涉及
	LogType *ShowAccessTopRequestLogType `json:"log_type,omitempty"`

	// **参数解释**： 聚合类型 **约束限制**： 不涉及 **取值范围**： strategy_hit_info 策略的命中趋势 strategy_dashboard 策略命中概览 top_deny_rule TOP命中拦截策略 dst_ip TOP拦截目的IP src_ip TOP拦截源IP dst_port TOP拦截端口 dst_region TOP拦截目的地区 src_region TOP拦截源地区 **默认取值**： 不涉及
	Item ShowAccessTopRequestItem `json:"item"`

	// **参数解释**： 规则ID列表 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	RuleId *[]string `json:"rule_id,omitempty"`
}

func (o ShowAccessTopRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAccessTopRequest struct{}"
	}

	return strings.Join([]string{"ShowAccessTopRequest", string(data)}, " ")
}

type ShowAccessTopRequestRange struct {
	value int32
}

type ShowAccessTopRequestRangeEnum struct {
	E_0 ShowAccessTopRequestRange
	E_1 ShowAccessTopRequestRange
	E_2 ShowAccessTopRequestRange
}

func GetShowAccessTopRequestRangeEnum() ShowAccessTopRequestRangeEnum {
	return ShowAccessTopRequestRangeEnum{
		E_0: ShowAccessTopRequestRange{
			value: 0,
		}, E_1: ShowAccessTopRequestRange{
			value: 1,
		}, E_2: ShowAccessTopRequestRange{
			value: 2,
		},
	}
}

func (c ShowAccessTopRequestRange) Value() int32 {
	return c.value
}

func (c ShowAccessTopRequestRange) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAccessTopRequestRange) UnmarshalJSON(b []byte) error {
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

type ShowAccessTopRequestDirection struct {
	value string
}

type ShowAccessTopRequestDirectionEnum struct {
	IN2OUT ShowAccessTopRequestDirection
	OUT2IN ShowAccessTopRequestDirection
}

func GetShowAccessTopRequestDirectionEnum() ShowAccessTopRequestDirectionEnum {
	return ShowAccessTopRequestDirectionEnum{
		IN2OUT: ShowAccessTopRequestDirection{
			value: "in2out",
		},
		OUT2IN: ShowAccessTopRequestDirection{
			value: "out2in",
		},
	}
}

func (c ShowAccessTopRequestDirection) Value() string {
	return c.value
}

func (c ShowAccessTopRequestDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAccessTopRequestDirection) UnmarshalJSON(b []byte) error {
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

type ShowAccessTopRequestLogType struct {
	value string
}

type ShowAccessTopRequestLogTypeEnum struct {
	INTERNET ShowAccessTopRequestLogType
	NAT      ShowAccessTopRequestLogType
	VPC      ShowAccessTopRequestLogType
	VGW      ShowAccessTopRequestLogType
}

func GetShowAccessTopRequestLogTypeEnum() ShowAccessTopRequestLogTypeEnum {
	return ShowAccessTopRequestLogTypeEnum{
		INTERNET: ShowAccessTopRequestLogType{
			value: "internet",
		},
		NAT: ShowAccessTopRequestLogType{
			value: "nat",
		},
		VPC: ShowAccessTopRequestLogType{
			value: "vpc",
		},
		VGW: ShowAccessTopRequestLogType{
			value: "vgw",
		},
	}
}

func (c ShowAccessTopRequestLogType) Value() string {
	return c.value
}

func (c ShowAccessTopRequestLogType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAccessTopRequestLogType) UnmarshalJSON(b []byte) error {
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

type ShowAccessTopRequestItem struct {
	value string
}

type ShowAccessTopRequestItemEnum struct {
	STRATEGY_DASHBOARD ShowAccessTopRequestItem
	STRATEGY_HIT_INFO  ShowAccessTopRequestItem
	TOP_DENY_RULE      ShowAccessTopRequestItem
	SRC_IP             ShowAccessTopRequestItem
	DST_IP             ShowAccessTopRequestItem
	SRC_REGION         ShowAccessTopRequestItem
	DST_REGION         ShowAccessTopRequestItem
	DST_PORT           ShowAccessTopRequestItem
}

func GetShowAccessTopRequestItemEnum() ShowAccessTopRequestItemEnum {
	return ShowAccessTopRequestItemEnum{
		STRATEGY_DASHBOARD: ShowAccessTopRequestItem{
			value: "strategy_dashboard",
		},
		STRATEGY_HIT_INFO: ShowAccessTopRequestItem{
			value: "strategy_hit_info",
		},
		TOP_DENY_RULE: ShowAccessTopRequestItem{
			value: "top_deny_rule",
		},
		SRC_IP: ShowAccessTopRequestItem{
			value: "src_ip",
		},
		DST_IP: ShowAccessTopRequestItem{
			value: "dst_ip",
		},
		SRC_REGION: ShowAccessTopRequestItem{
			value: "src_region",
		},
		DST_REGION: ShowAccessTopRequestItem{
			value: "dst_region",
		},
		DST_PORT: ShowAccessTopRequestItem{
			value: "dst_port",
		},
	}
}

func (c ShowAccessTopRequestItem) Value() string {
	return c.value
}

func (c ShowAccessTopRequestItem) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAccessTopRequestItem) UnmarshalJSON(b []byte) error {
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
