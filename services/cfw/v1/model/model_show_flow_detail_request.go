package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowFlowDetailRequest Request Object
type ShowFlowDetailRequest struct {

	// **参数解释**： 防火墙ID，用户创建防火墙实例后产生的唯一ID，配置后可区分不同防火墙，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	FwInstanceId string `json:"fw_instance_id"`

	// **参数解释**： 时间范围  **约束限制**： 不涉及 **取值范围**： 0为近一时 1近一天 2近七天   **默认取值**： 不涉及
	Range *ShowFlowDetailRequestRange `json:"range,omitempty"`

	// **参数解释**： 日志类型 **约束限制**： 不涉及 **取值范围**： internet为南北向日志、nat为nat场景日志，vpc为东西向日志，vgw为vgw场景日志 **默认取值**： 不涉及
	LogType ShowFlowDetailRequestLogType `json:"log_type"`

	// **参数解释**： 会话方向 **约束限制**： 不涉及 **取值范围**： in2out为出云方向 out2in为入云方向 **默认取值**： 不涉及
	Direction *ShowFlowDetailRequestDirection `json:"direction,omitempty"`

	// **参数解释**： 开始时间 **约束限制**： 不涉及 **取值范围**： 毫秒级时间戳 **默认取值**： 不涉及
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**： 结束时间 **约束限制**： 不涉及 **取值范围**： 毫秒级时间戳 **默认取值**： 不涉及
	EndTime *int64 `json:"end_time,omitempty"`

	// **参数解释**： VGW ID **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	VgwId *[]string `json:"vgw_id,omitempty"`

	// **参数解释**： IP类型 **约束限制**： 不涉及 **取值范围**： public 公网IP private 私网IP open_port **默认取值**： 不涉及
	AssetType *ShowFlowDetailRequestAssetType `json:"asset_type,omitempty"`

	// **参数解释**： 聚合类型 **约束限制**： 不涉及 **取值范围**： src_ip 源IP dst_ip 目的IP dst_port 目的端口 protocol　协议 dst_host　目的域名 app　应用 dst_region_name　目的地区 src_region_name　源地区 **默认取值**： 不涉及
	Item ShowFlowDetailRequestItem `json:"item"`

	// **参数解释**： 统计对象 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Value string `json:"value"`
}

func (o ShowFlowDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlowDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowFlowDetailRequest", string(data)}, " ")
}

type ShowFlowDetailRequestRange struct {
	value int32
}

type ShowFlowDetailRequestRangeEnum struct {
	E_0 ShowFlowDetailRequestRange
	E_1 ShowFlowDetailRequestRange
	E_2 ShowFlowDetailRequestRange
}

func GetShowFlowDetailRequestRangeEnum() ShowFlowDetailRequestRangeEnum {
	return ShowFlowDetailRequestRangeEnum{
		E_0: ShowFlowDetailRequestRange{
			value: 0,
		}, E_1: ShowFlowDetailRequestRange{
			value: 1,
		}, E_2: ShowFlowDetailRequestRange{
			value: 2,
		},
	}
}

func (c ShowFlowDetailRequestRange) Value() int32 {
	return c.value
}

func (c ShowFlowDetailRequestRange) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFlowDetailRequestRange) UnmarshalJSON(b []byte) error {
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

type ShowFlowDetailRequestLogType struct {
	value string
}

type ShowFlowDetailRequestLogTypeEnum struct {
	INTERNET ShowFlowDetailRequestLogType
	NAT      ShowFlowDetailRequestLogType
	VPC      ShowFlowDetailRequestLogType
	VGW      ShowFlowDetailRequestLogType
}

func GetShowFlowDetailRequestLogTypeEnum() ShowFlowDetailRequestLogTypeEnum {
	return ShowFlowDetailRequestLogTypeEnum{
		INTERNET: ShowFlowDetailRequestLogType{
			value: "internet",
		},
		NAT: ShowFlowDetailRequestLogType{
			value: "nat",
		},
		VPC: ShowFlowDetailRequestLogType{
			value: "vpc",
		},
		VGW: ShowFlowDetailRequestLogType{
			value: "vgw",
		},
	}
}

func (c ShowFlowDetailRequestLogType) Value() string {
	return c.value
}

func (c ShowFlowDetailRequestLogType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFlowDetailRequestLogType) UnmarshalJSON(b []byte) error {
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

type ShowFlowDetailRequestDirection struct {
	value string
}

type ShowFlowDetailRequestDirectionEnum struct {
	IN2OUT ShowFlowDetailRequestDirection
	OUT2IN ShowFlowDetailRequestDirection
}

func GetShowFlowDetailRequestDirectionEnum() ShowFlowDetailRequestDirectionEnum {
	return ShowFlowDetailRequestDirectionEnum{
		IN2OUT: ShowFlowDetailRequestDirection{
			value: "in2out",
		},
		OUT2IN: ShowFlowDetailRequestDirection{
			value: "out2in",
		},
	}
}

func (c ShowFlowDetailRequestDirection) Value() string {
	return c.value
}

func (c ShowFlowDetailRequestDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFlowDetailRequestDirection) UnmarshalJSON(b []byte) error {
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

type ShowFlowDetailRequestAssetType struct {
	value string
}

type ShowFlowDetailRequestAssetTypeEnum struct {
	PUBLIC  ShowFlowDetailRequestAssetType
	PRIVATE ShowFlowDetailRequestAssetType
}

func GetShowFlowDetailRequestAssetTypeEnum() ShowFlowDetailRequestAssetTypeEnum {
	return ShowFlowDetailRequestAssetTypeEnum{
		PUBLIC: ShowFlowDetailRequestAssetType{
			value: "public",
		},
		PRIVATE: ShowFlowDetailRequestAssetType{
			value: "private",
		},
	}
}

func (c ShowFlowDetailRequestAssetType) Value() string {
	return c.value
}

func (c ShowFlowDetailRequestAssetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFlowDetailRequestAssetType) UnmarshalJSON(b []byte) error {
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

type ShowFlowDetailRequestItem struct {
	value string
}

type ShowFlowDetailRequestItemEnum struct {
	SRC_IP          ShowFlowDetailRequestItem
	DST_IP          ShowFlowDetailRequestItem
	DST_PORT        ShowFlowDetailRequestItem
	PROTOCOL        ShowFlowDetailRequestItem
	DST_HOST        ShowFlowDetailRequestItem
	APP             ShowFlowDetailRequestItem
	DST_REGION_NAME ShowFlowDetailRequestItem
	SRC_REGION_NAME ShowFlowDetailRequestItem
}

func GetShowFlowDetailRequestItemEnum() ShowFlowDetailRequestItemEnum {
	return ShowFlowDetailRequestItemEnum{
		SRC_IP: ShowFlowDetailRequestItem{
			value: "src_ip",
		},
		DST_IP: ShowFlowDetailRequestItem{
			value: "dst_ip",
		},
		DST_PORT: ShowFlowDetailRequestItem{
			value: "dst_port",
		},
		PROTOCOL: ShowFlowDetailRequestItem{
			value: "protocol",
		},
		DST_HOST: ShowFlowDetailRequestItem{
			value: "dst_host",
		},
		APP: ShowFlowDetailRequestItem{
			value: "app",
		},
		DST_REGION_NAME: ShowFlowDetailRequestItem{
			value: "dst_region_name",
		},
		SRC_REGION_NAME: ShowFlowDetailRequestItem{
			value: "src_region_name",
		},
	}
}

func (c ShowFlowDetailRequestItem) Value() string {
	return c.value
}

func (c ShowFlowDetailRequestItem) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFlowDetailRequestItem) UnmarshalJSON(b []byte) error {
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
