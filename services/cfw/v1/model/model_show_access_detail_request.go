package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAccessDetailRequest Request Object
type ShowAccessDetailRequest struct {

	// **参数解释**： 防火墙ID，用户创建防火墙实例后产生的唯一ID，配置后可区分不同防火墙，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	FwInstanceId string `json:"fw_instance_id"`

	// **参数解释**： 时间范围  **约束限制**： 不涉及 **取值范围**： 0为近一时 1近一天 2近七天   **默认取值**： 不涉及
	Range *ShowAccessDetailRequestRange `json:"range,omitempty"`

	// **参数解释**： 会话方向 **约束限制**： 不涉及 **取值范围**： in2out为出云方向 out2in为入云方向 **默认取值**： 不涉及
	Direction *ShowAccessDetailRequestDirection `json:"direction,omitempty"`

	// **参数解释**： 开始时间 **约束限制**： 不涉及 **取值范围**： 毫秒级时间戳 **默认取值**： 不涉及
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**： 结束时间 **约束限制**： 不涉及 **取值范围**： 毫秒级时间戳 **默认取值**： 不涉及
	EndTime *int64 `json:"end_time,omitempty"`

	// **参数解释**： VGW ID **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	VgwId *[]string `json:"vgw_id,omitempty"`

	// **参数解释**： 日志类型 **约束限制**： 不涉及 **取值范围**： internet为南北向日志、nat为nat场景日志，vpc为东西向日志，vgw为vgw场景日志 **默认取值**： 不涉及
	LogType *ShowAccessDetailRequestLogType `json:"log_type,omitempty"`

	// **参数解释**： 聚合类型 **约束限制**： 不涉及 **取值范围**： top_deny_rule 命中拦截策略 dst_ip 拦截目的IP src_ip 拦截源IP dst_port 拦截端口 dst_region 拦截目的地区 src_region 拦截源地区 **默认取值**： 不涉及
	Item ShowAccessDetailRequestItem `json:"item"`

	// **参数解释**： 聚合对象 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ItemId string `json:"item_id"`
}

func (o ShowAccessDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAccessDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowAccessDetailRequest", string(data)}, " ")
}

type ShowAccessDetailRequestRange struct {
	value int32
}

type ShowAccessDetailRequestRangeEnum struct {
	E_0 ShowAccessDetailRequestRange
	E_1 ShowAccessDetailRequestRange
	E_2 ShowAccessDetailRequestRange
}

func GetShowAccessDetailRequestRangeEnum() ShowAccessDetailRequestRangeEnum {
	return ShowAccessDetailRequestRangeEnum{
		E_0: ShowAccessDetailRequestRange{
			value: 0,
		}, E_1: ShowAccessDetailRequestRange{
			value: 1,
		}, E_2: ShowAccessDetailRequestRange{
			value: 2,
		},
	}
}

func (c ShowAccessDetailRequestRange) Value() int32 {
	return c.value
}

func (c ShowAccessDetailRequestRange) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAccessDetailRequestRange) UnmarshalJSON(b []byte) error {
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

type ShowAccessDetailRequestDirection struct {
	value string
}

type ShowAccessDetailRequestDirectionEnum struct {
	IN2OUT ShowAccessDetailRequestDirection
	OUT2IN ShowAccessDetailRequestDirection
}

func GetShowAccessDetailRequestDirectionEnum() ShowAccessDetailRequestDirectionEnum {
	return ShowAccessDetailRequestDirectionEnum{
		IN2OUT: ShowAccessDetailRequestDirection{
			value: "in2out",
		},
		OUT2IN: ShowAccessDetailRequestDirection{
			value: "out2in",
		},
	}
}

func (c ShowAccessDetailRequestDirection) Value() string {
	return c.value
}

func (c ShowAccessDetailRequestDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAccessDetailRequestDirection) UnmarshalJSON(b []byte) error {
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

type ShowAccessDetailRequestLogType struct {
	value string
}

type ShowAccessDetailRequestLogTypeEnum struct {
	INTERNET ShowAccessDetailRequestLogType
	NAT      ShowAccessDetailRequestLogType
	VPC      ShowAccessDetailRequestLogType
	VGW      ShowAccessDetailRequestLogType
}

func GetShowAccessDetailRequestLogTypeEnum() ShowAccessDetailRequestLogTypeEnum {
	return ShowAccessDetailRequestLogTypeEnum{
		INTERNET: ShowAccessDetailRequestLogType{
			value: "internet",
		},
		NAT: ShowAccessDetailRequestLogType{
			value: "nat",
		},
		VPC: ShowAccessDetailRequestLogType{
			value: "vpc",
		},
		VGW: ShowAccessDetailRequestLogType{
			value: "vgw",
		},
	}
}

func (c ShowAccessDetailRequestLogType) Value() string {
	return c.value
}

func (c ShowAccessDetailRequestLogType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAccessDetailRequestLogType) UnmarshalJSON(b []byte) error {
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

type ShowAccessDetailRequestItem struct {
	value string
}

type ShowAccessDetailRequestItemEnum struct {
	TOP_DENY_RULE ShowAccessDetailRequestItem
	SRC_IP        ShowAccessDetailRequestItem
	DST_IP        ShowAccessDetailRequestItem
	SRC_REGION    ShowAccessDetailRequestItem
	DST_REGION    ShowAccessDetailRequestItem
	DST_PORT      ShowAccessDetailRequestItem
}

func GetShowAccessDetailRequestItemEnum() ShowAccessDetailRequestItemEnum {
	return ShowAccessDetailRequestItemEnum{
		TOP_DENY_RULE: ShowAccessDetailRequestItem{
			value: "top_deny_rule",
		},
		SRC_IP: ShowAccessDetailRequestItem{
			value: "src_ip",
		},
		DST_IP: ShowAccessDetailRequestItem{
			value: "dst_ip",
		},
		SRC_REGION: ShowAccessDetailRequestItem{
			value: "src_region",
		},
		DST_REGION: ShowAccessDetailRequestItem{
			value: "dst_region",
		},
		DST_PORT: ShowAccessDetailRequestItem{
			value: "dst_port",
		},
	}
}

func (c ShowAccessDetailRequestItem) Value() string {
	return c.value
}

func (c ShowAccessDetailRequestItem) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAccessDetailRequestItem) UnmarshalJSON(b []byte) error {
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
