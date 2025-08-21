package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowFlowTopRequest Request Object
type ShowFlowTopRequest struct {

	// **参数解释**： 防火墙ID，用户创建防火墙实例后产生的唯一ID，配置后可区分不同防火墙，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	FwInstanceId string `json:"fw_instance_id"`

	// **参数解释**： 时间范围  **约束限制**： 不涉及 **取值范围**： 0为近一时 1近一天 2近七天   **默认取值**： 不涉及
	Range *ShowFlowTopRequestRange `json:"range,omitempty"`

	// **参数解释**： 日志类型 **约束限制**： 不涉及 **取值范围**： internet为南北向日志、nat为nat场景日志，vpc为东西向日志，vgw为vgw场景日志 **默认取值**： 不涉及
	LogType ShowFlowTopRequestLogType `json:"log_type"`

	// **参数解释**： 会话方向 **约束限制**： 不涉及 **取值范围**： in2out为出云方向 out2in为入云方向 **默认取值**： 不涉及
	Direction *ShowFlowTopRequestDirection `json:"direction,omitempty"`

	// **参数解释**： 开始时间 **约束限制**： 不涉及 **取值范围**： 毫秒级时间戳 **默认取值**： 不涉及
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**： 结束时间 **约束限制**： 不涉及 **取值范围**： 毫秒级时间戳 **默认取值**： 不涉及
	EndTime *int64 `json:"end_time,omitempty"`

	// **参数解释**： VGW ID **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	VgwId *[]string `json:"vgw_id,omitempty"`

	// **参数解释**： IP类型 **约束限制**： 不涉及 **取值范围**： public 公网IP private 私网IP open_port **默认取值**： 不涉及
	AssetType *ShowFlowTopRequestAssetType `json:"asset_type,omitempty"`

	// **参数解释**： 聚合类型 **约束限制**： 不涉及 **取值范围**： src_ip 源IP dst_ip 目的IP dst_port 目的端口 protocol　协议 dst_host　目的域名 app　应用 dst_region_name　目的地区 src_region_name　源地区 **默认取值**： 不涉及
	Item []string `json:"item"`

	// **参数解释**： 排序条件 **约束限制**： 不涉及 **取值范围**： record 会话条数 byte 会话流量 **默认取值**： 5
	Order ShowFlowTopRequestOrder `json:"order"`

	// **参数解释**： 聚合条数 **约束限制**： 不涉及 **取值范围**： 0到10条 **默认取值**： 5
	Size *int32 `json:"size,omitempty"`
}

func (o ShowFlowTopRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlowTopRequest struct{}"
	}

	return strings.Join([]string{"ShowFlowTopRequest", string(data)}, " ")
}

type ShowFlowTopRequestRange struct {
	value int32
}

type ShowFlowTopRequestRangeEnum struct {
	E_0 ShowFlowTopRequestRange
	E_1 ShowFlowTopRequestRange
	E_2 ShowFlowTopRequestRange
}

func GetShowFlowTopRequestRangeEnum() ShowFlowTopRequestRangeEnum {
	return ShowFlowTopRequestRangeEnum{
		E_0: ShowFlowTopRequestRange{
			value: 0,
		}, E_1: ShowFlowTopRequestRange{
			value: 1,
		}, E_2: ShowFlowTopRequestRange{
			value: 2,
		},
	}
}

func (c ShowFlowTopRequestRange) Value() int32 {
	return c.value
}

func (c ShowFlowTopRequestRange) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFlowTopRequestRange) UnmarshalJSON(b []byte) error {
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

type ShowFlowTopRequestLogType struct {
	value string
}

type ShowFlowTopRequestLogTypeEnum struct {
	INTERNET ShowFlowTopRequestLogType
	NAT      ShowFlowTopRequestLogType
	VPC      ShowFlowTopRequestLogType
	VGW      ShowFlowTopRequestLogType
}

func GetShowFlowTopRequestLogTypeEnum() ShowFlowTopRequestLogTypeEnum {
	return ShowFlowTopRequestLogTypeEnum{
		INTERNET: ShowFlowTopRequestLogType{
			value: "internet",
		},
		NAT: ShowFlowTopRequestLogType{
			value: "nat",
		},
		VPC: ShowFlowTopRequestLogType{
			value: "vpc",
		},
		VGW: ShowFlowTopRequestLogType{
			value: "vgw",
		},
	}
}

func (c ShowFlowTopRequestLogType) Value() string {
	return c.value
}

func (c ShowFlowTopRequestLogType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFlowTopRequestLogType) UnmarshalJSON(b []byte) error {
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

type ShowFlowTopRequestDirection struct {
	value string
}

type ShowFlowTopRequestDirectionEnum struct {
	IN2OUT ShowFlowTopRequestDirection
	OUT2IN ShowFlowTopRequestDirection
}

func GetShowFlowTopRequestDirectionEnum() ShowFlowTopRequestDirectionEnum {
	return ShowFlowTopRequestDirectionEnum{
		IN2OUT: ShowFlowTopRequestDirection{
			value: "in2out",
		},
		OUT2IN: ShowFlowTopRequestDirection{
			value: "out2in",
		},
	}
}

func (c ShowFlowTopRequestDirection) Value() string {
	return c.value
}

func (c ShowFlowTopRequestDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFlowTopRequestDirection) UnmarshalJSON(b []byte) error {
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

type ShowFlowTopRequestAssetType struct {
	value string
}

type ShowFlowTopRequestAssetTypeEnum struct {
	PUBLIC  ShowFlowTopRequestAssetType
	PRIVATE ShowFlowTopRequestAssetType
}

func GetShowFlowTopRequestAssetTypeEnum() ShowFlowTopRequestAssetTypeEnum {
	return ShowFlowTopRequestAssetTypeEnum{
		PUBLIC: ShowFlowTopRequestAssetType{
			value: "public",
		},
		PRIVATE: ShowFlowTopRequestAssetType{
			value: "private",
		},
	}
}

func (c ShowFlowTopRequestAssetType) Value() string {
	return c.value
}

func (c ShowFlowTopRequestAssetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFlowTopRequestAssetType) UnmarshalJSON(b []byte) error {
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

type ShowFlowTopRequestOrder struct {
	value string
}

type ShowFlowTopRequestOrderEnum struct {
	RECORD ShowFlowTopRequestOrder
	BYTE   ShowFlowTopRequestOrder
}

func GetShowFlowTopRequestOrderEnum() ShowFlowTopRequestOrderEnum {
	return ShowFlowTopRequestOrderEnum{
		RECORD: ShowFlowTopRequestOrder{
			value: "record",
		},
		BYTE: ShowFlowTopRequestOrder{
			value: "byte",
		},
	}
}

func (c ShowFlowTopRequestOrder) Value() string {
	return c.value
}

func (c ShowFlowTopRequestOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFlowTopRequestOrder) UnmarshalJSON(b []byte) error {
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
