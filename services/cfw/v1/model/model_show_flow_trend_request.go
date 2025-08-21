package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowFlowTrendRequest Request Object
type ShowFlowTrendRequest struct {

	// **参数解释**： 防火墙ID，用户创建防火墙实例后产生的唯一ID，配置后可区分不同防火墙，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	FwInstanceId string `json:"fw_instance_id"`

	// **参数解释**： 时间范围  **约束限制**： 不涉及 **取值范围**： 0为近一时 1近一天 2近七天   **默认取值**： 不涉及
	Range *ShowFlowTrendRequestRange `json:"range,omitempty"`

	// **参数解释**： 日志类型 **约束限制**： 不涉及 **取值范围**： internet为南北向日志、nat为nat场景日志，vpc为东西向日志，vgw为vgw场景日志 **默认取值**： 不涉及
	LogType ShowFlowTrendRequestLogType `json:"log_type"`

	// **参数解释**： 会话方向 **约束限制**： 不涉及 **取值范围**： in2out为出云方向 out2in为入云方向 **默认取值**： 不涉及
	Direction *ShowFlowTrendRequestDirection `json:"direction,omitempty"`

	// **参数解释**： 开始时间 **约束限制**： 不涉及 **取值范围**： 毫秒级时间戳 **默认取值**： 不涉及
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**： 结束时间 **约束限制**： 不涉及 **取值范围**： 毫秒级时间戳 **默认取值**： 不涉及
	EndTime *int64 `json:"end_time,omitempty"`

	// **参数解释**： VGW ID **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	VgwId *[]string `json:"vgw_id,omitempty"`

	// **参数解释**： IP类型 **约束限制**： 不涉及 **取值范围**： public 公网IP private 私网IP open_port **默认取值**： 不涉及
	AssetType *ShowFlowTrendRequestAssetType `json:"asset_type,omitempty"`

	// ips
	Ip *[]string `json:"ip,omitempty"`

	// vpcs
	Vpc *[]string `json:"vpc,omitempty"`
}

func (o ShowFlowTrendRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlowTrendRequest struct{}"
	}

	return strings.Join([]string{"ShowFlowTrendRequest", string(data)}, " ")
}

type ShowFlowTrendRequestRange struct {
	value int32
}

type ShowFlowTrendRequestRangeEnum struct {
	E_0 ShowFlowTrendRequestRange
	E_1 ShowFlowTrendRequestRange
	E_2 ShowFlowTrendRequestRange
}

func GetShowFlowTrendRequestRangeEnum() ShowFlowTrendRequestRangeEnum {
	return ShowFlowTrendRequestRangeEnum{
		E_0: ShowFlowTrendRequestRange{
			value: 0,
		}, E_1: ShowFlowTrendRequestRange{
			value: 1,
		}, E_2: ShowFlowTrendRequestRange{
			value: 2,
		},
	}
}

func (c ShowFlowTrendRequestRange) Value() int32 {
	return c.value
}

func (c ShowFlowTrendRequestRange) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFlowTrendRequestRange) UnmarshalJSON(b []byte) error {
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

type ShowFlowTrendRequestLogType struct {
	value string
}

type ShowFlowTrendRequestLogTypeEnum struct {
	INTERNET ShowFlowTrendRequestLogType
	NAT      ShowFlowTrendRequestLogType
	VPC      ShowFlowTrendRequestLogType
	VGW      ShowFlowTrendRequestLogType
}

func GetShowFlowTrendRequestLogTypeEnum() ShowFlowTrendRequestLogTypeEnum {
	return ShowFlowTrendRequestLogTypeEnum{
		INTERNET: ShowFlowTrendRequestLogType{
			value: "internet",
		},
		NAT: ShowFlowTrendRequestLogType{
			value: "nat",
		},
		VPC: ShowFlowTrendRequestLogType{
			value: "vpc",
		},
		VGW: ShowFlowTrendRequestLogType{
			value: "vgw",
		},
	}
}

func (c ShowFlowTrendRequestLogType) Value() string {
	return c.value
}

func (c ShowFlowTrendRequestLogType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFlowTrendRequestLogType) UnmarshalJSON(b []byte) error {
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

type ShowFlowTrendRequestDirection struct {
	value string
}

type ShowFlowTrendRequestDirectionEnum struct {
	IN2OUT ShowFlowTrendRequestDirection
	OUT2IN ShowFlowTrendRequestDirection
}

func GetShowFlowTrendRequestDirectionEnum() ShowFlowTrendRequestDirectionEnum {
	return ShowFlowTrendRequestDirectionEnum{
		IN2OUT: ShowFlowTrendRequestDirection{
			value: "in2out",
		},
		OUT2IN: ShowFlowTrendRequestDirection{
			value: "out2in",
		},
	}
}

func (c ShowFlowTrendRequestDirection) Value() string {
	return c.value
}

func (c ShowFlowTrendRequestDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFlowTrendRequestDirection) UnmarshalJSON(b []byte) error {
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

type ShowFlowTrendRequestAssetType struct {
	value string
}

type ShowFlowTrendRequestAssetTypeEnum struct {
	PUBLIC  ShowFlowTrendRequestAssetType
	PRIVATE ShowFlowTrendRequestAssetType
}

func GetShowFlowTrendRequestAssetTypeEnum() ShowFlowTrendRequestAssetTypeEnum {
	return ShowFlowTrendRequestAssetTypeEnum{
		PUBLIC: ShowFlowTrendRequestAssetType{
			value: "public",
		},
		PRIVATE: ShowFlowTrendRequestAssetType{
			value: "private",
		},
	}
}

func (c ShowFlowTrendRequestAssetType) Value() string {
	return c.value
}

func (c ShowFlowTrendRequestAssetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFlowTrendRequestAssetType) UnmarshalJSON(b []byte) error {
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
