package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowLogsCountRequest Request Object
type ShowLogsCountRequest struct {

	// **参数解释**： 防火墙ID，用户创建防火墙实例后产生的唯一ID，配置后可区分不同防火墙，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	FwInstanceId string `json:"fw_instance_id"`

	// **参数解释**： 时间范围  **约束限制**： 不涉及 **取值范围**： 0为近一时 1近一天 2近七天   **默认取值**： 不涉及
	Range *ShowLogsCountRequestRange `json:"range,omitempty"`

	// **参数解释**： 开始时间 **约束限制**： 不涉及 **取值范围**： 毫秒级时间戳 **默认取值**： 不涉及
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**： 结束时间 **约束限制**： 不涉及 **取值范围**： 毫秒级时间戳 **默认取值**： 不涉及
	EndTime *int64 `json:"end_time,omitempty"`

	// **参数解释**： VGW ID **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	VgwId *[]string `json:"vgw_id,omitempty"`

	// **参数解释**： 聚合类型 **约束限制**： 不涉及 **取值范围**： risk_ip 访问风险IP数量 risk_host 访问风险域名数量  open_port 开放端口的数量 **默认取值**： 不涉及
	Item ShowLogsCountRequestItem `json:"item"`
}

func (o ShowLogsCountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLogsCountRequest struct{}"
	}

	return strings.Join([]string{"ShowLogsCountRequest", string(data)}, " ")
}

type ShowLogsCountRequestRange struct {
	value int32
}

type ShowLogsCountRequestRangeEnum struct {
	E_0 ShowLogsCountRequestRange
	E_1 ShowLogsCountRequestRange
	E_2 ShowLogsCountRequestRange
}

func GetShowLogsCountRequestRangeEnum() ShowLogsCountRequestRangeEnum {
	return ShowLogsCountRequestRangeEnum{
		E_0: ShowLogsCountRequestRange{
			value: 0,
		}, E_1: ShowLogsCountRequestRange{
			value: 1,
		}, E_2: ShowLogsCountRequestRange{
			value: 2,
		},
	}
}

func (c ShowLogsCountRequestRange) Value() int32 {
	return c.value
}

func (c ShowLogsCountRequestRange) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowLogsCountRequestRange) UnmarshalJSON(b []byte) error {
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

type ShowLogsCountRequestItem struct {
	value string
}

type ShowLogsCountRequestItemEnum struct {
	RISK_IP   ShowLogsCountRequestItem
	RISK_HOST ShowLogsCountRequestItem
	OPEN_PORT ShowLogsCountRequestItem
}

func GetShowLogsCountRequestItemEnum() ShowLogsCountRequestItemEnum {
	return ShowLogsCountRequestItemEnum{
		RISK_IP: ShowLogsCountRequestItem{
			value: "risk_ip",
		},
		RISK_HOST: ShowLogsCountRequestItem{
			value: "risk_host",
		},
		OPEN_PORT: ShowLogsCountRequestItem{
			value: "open_port",
		},
	}
}

func (c ShowLogsCountRequestItem) Value() string {
	return c.value
}

func (c ShowLogsCountRequestItem) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowLogsCountRequestItem) UnmarshalJSON(b []byte) error {
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
