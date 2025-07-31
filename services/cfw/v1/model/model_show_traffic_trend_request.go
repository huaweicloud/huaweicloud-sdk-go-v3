package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowTrafficTrendRequest Request Object
type ShowTrafficTrendRequest struct {

	// **参数解释**： 防火墙ID，用户创建防火墙实例后产生的唯一ID，配置后可区分不同防火墙，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	FwInstanceId string `json:"fw_instance_id"`

	// **参数解释**： 时间范围  **约束限制**： 不涉及 **取值范围**： 0为近一时 1近一天 2近七天   **默认取值**： 不涉及
	Range *ShowTrafficTrendRequestRange `json:"range,omitempty"`

	// **参数解释**： 日志类型 **约束限制**： 不涉及 **取值范围**： internet为南北向日志、nat为nat场景日志，vpc为东西向日志，vgw为vgw场景日志 **默认取值**： 不涉及
	LogType ShowTrafficTrendRequestLogType `json:"log_type"`

	// **参数解释**： 开始时间 **约束限制**： 不涉及 **取值范围**： 毫秒级时间戳 **默认取值**： 不涉及
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**： 结束时间 **约束限制**： 不涉及 **取值范围**： 毫秒级时间戳 **默认取值**： 不涉及
	EndTime *int64 `json:"end_time,omitempty"`

	// **参数解释**： VGW ID **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	VgwId *[]string `json:"vgw_id,omitempty"`

	// **参数解释**： 聚合类型 **约束限制**： 不涉及 **取值范围**： avg为平均值 max为最大值 **默认取值**： 不涉及
	AggType ShowTrafficTrendRequestAggType `json:"agg_type"`

	// **参数解释**： IP地址列表，查看IP的流量趋势 **约束限制**： 不涉及 **取值范围**： IP地址 **默认取值**： 不涉及
	Ip *[]string `json:"ip,omitempty"`
}

func (o ShowTrafficTrendRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTrafficTrendRequest struct{}"
	}

	return strings.Join([]string{"ShowTrafficTrendRequest", string(data)}, " ")
}

type ShowTrafficTrendRequestRange struct {
	value int32
}

type ShowTrafficTrendRequestRangeEnum struct {
	E_0 ShowTrafficTrendRequestRange
	E_1 ShowTrafficTrendRequestRange
	E_2 ShowTrafficTrendRequestRange
}

func GetShowTrafficTrendRequestRangeEnum() ShowTrafficTrendRequestRangeEnum {
	return ShowTrafficTrendRequestRangeEnum{
		E_0: ShowTrafficTrendRequestRange{
			value: 0,
		}, E_1: ShowTrafficTrendRequestRange{
			value: 1,
		}, E_2: ShowTrafficTrendRequestRange{
			value: 2,
		},
	}
}

func (c ShowTrafficTrendRequestRange) Value() int32 {
	return c.value
}

func (c ShowTrafficTrendRequestRange) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTrafficTrendRequestRange) UnmarshalJSON(b []byte) error {
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

type ShowTrafficTrendRequestLogType struct {
	value string
}

type ShowTrafficTrendRequestLogTypeEnum struct {
	INTERNET ShowTrafficTrendRequestLogType
	NAT      ShowTrafficTrendRequestLogType
	VPC      ShowTrafficTrendRequestLogType
	VGW      ShowTrafficTrendRequestLogType
}

func GetShowTrafficTrendRequestLogTypeEnum() ShowTrafficTrendRequestLogTypeEnum {
	return ShowTrafficTrendRequestLogTypeEnum{
		INTERNET: ShowTrafficTrendRequestLogType{
			value: "internet",
		},
		NAT: ShowTrafficTrendRequestLogType{
			value: "nat",
		},
		VPC: ShowTrafficTrendRequestLogType{
			value: "vpc",
		},
		VGW: ShowTrafficTrendRequestLogType{
			value: "vgw",
		},
	}
}

func (c ShowTrafficTrendRequestLogType) Value() string {
	return c.value
}

func (c ShowTrafficTrendRequestLogType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTrafficTrendRequestLogType) UnmarshalJSON(b []byte) error {
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

type ShowTrafficTrendRequestAggType struct {
	value string
}

type ShowTrafficTrendRequestAggTypeEnum struct {
	AVG ShowTrafficTrendRequestAggType
	MAX ShowTrafficTrendRequestAggType
}

func GetShowTrafficTrendRequestAggTypeEnum() ShowTrafficTrendRequestAggTypeEnum {
	return ShowTrafficTrendRequestAggTypeEnum{
		AVG: ShowTrafficTrendRequestAggType{
			value: "avg",
		},
		MAX: ShowTrafficTrendRequestAggType{
			value: "max",
		},
	}
}

func (c ShowTrafficTrendRequestAggType) Value() string {
	return c.value
}

func (c ShowTrafficTrendRequestAggType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTrafficTrendRequestAggType) UnmarshalJSON(b []byte) error {
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
