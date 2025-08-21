package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListFlowStatisticRequest Request Object
type ListFlowStatisticRequest struct {

	// **参数解释**： 防火墙ID，用户创建防火墙实例后产生的唯一ID，配置后可区分不同防火墙，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	FwInstanceId string `json:"fw_instance_id"`

	// **参数解释**： 时间范围  **约束限制**： 不涉及 **取值范围**： 0为近一时 1近一天 2近七天   **默认取值**： 不涉及
	Range *ListFlowStatisticRequestRange `json:"range,omitempty"`

	// **参数解释**： 日志类型 **约束限制**： 不涉及 **取值范围**： internet为南北向日志、nat为nat场景日志，vpc为东西向日志，vgw为vgw场景日志 **默认取值**： 不涉及
	LogType ListFlowStatisticRequestLogType `json:"log_type"`

	// **参数解释**： 会话方向 **约束限制**： 不涉及 **取值范围**： in2out为出云方向 out2in为入云方向 **默认取值**： 不涉及
	Direction *ListFlowStatisticRequestDirection `json:"direction,omitempty"`

	// **参数解释**： 开始时间 **约束限制**： 不涉及 **取值范围**： 毫秒级时间戳 **默认取值**： 不涉及
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**： 结束时间 **约束限制**： 不涉及 **取值范围**： 毫秒级时间戳 **默认取值**： 不涉及
	EndTime *int64 `json:"end_time,omitempty"`

	// **参数解释**： VGW ID **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	VgwId *[]string `json:"vgw_id,omitempty"`

	// **参数解释**： IP类型 **约束限制**： 不涉及 **取值范围**： public 公网IP private 私网IP open_port **默认取值**： 不涉及
	AssetType *ListFlowStatisticRequestAssetType `json:"asset_type,omitempty"`

	// **参数解释**： 聚合类型 **约束限制**： 不涉及 **取值范围**： src_ip 源IP dst_ip 目的IP dst_port 目的端口 protocol　协议 dst_host　目的域名 app　应用 dst_region_name　目的地区 src_region_name　源地区 risk_ip 风险IP risk_host 风险域名 open_port 开放端口 **默认取值**： 不涉及
	Item ListFlowStatisticRequestItem `json:"item"`

	// **参数解释**： 聚合条数 **约束限制**： 不涉及 **取值范围**： 0到10条 **默认取值**： 5
	Size *int32 `json:"size,omitempty"`
}

func (o ListFlowStatisticRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlowStatisticRequest struct{}"
	}

	return strings.Join([]string{"ListFlowStatisticRequest", string(data)}, " ")
}

type ListFlowStatisticRequestRange struct {
	value int32
}

type ListFlowStatisticRequestRangeEnum struct {
	E_0 ListFlowStatisticRequestRange
	E_1 ListFlowStatisticRequestRange
	E_2 ListFlowStatisticRequestRange
}

func GetListFlowStatisticRequestRangeEnum() ListFlowStatisticRequestRangeEnum {
	return ListFlowStatisticRequestRangeEnum{
		E_0: ListFlowStatisticRequestRange{
			value: 0,
		}, E_1: ListFlowStatisticRequestRange{
			value: 1,
		}, E_2: ListFlowStatisticRequestRange{
			value: 2,
		},
	}
}

func (c ListFlowStatisticRequestRange) Value() int32 {
	return c.value
}

func (c ListFlowStatisticRequestRange) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFlowStatisticRequestRange) UnmarshalJSON(b []byte) error {
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

type ListFlowStatisticRequestLogType struct {
	value string
}

type ListFlowStatisticRequestLogTypeEnum struct {
	INTERNET ListFlowStatisticRequestLogType
	NAT      ListFlowStatisticRequestLogType
	VPC      ListFlowStatisticRequestLogType
	VGW      ListFlowStatisticRequestLogType
}

func GetListFlowStatisticRequestLogTypeEnum() ListFlowStatisticRequestLogTypeEnum {
	return ListFlowStatisticRequestLogTypeEnum{
		INTERNET: ListFlowStatisticRequestLogType{
			value: "internet",
		},
		NAT: ListFlowStatisticRequestLogType{
			value: "nat",
		},
		VPC: ListFlowStatisticRequestLogType{
			value: "vpc",
		},
		VGW: ListFlowStatisticRequestLogType{
			value: "vgw",
		},
	}
}

func (c ListFlowStatisticRequestLogType) Value() string {
	return c.value
}

func (c ListFlowStatisticRequestLogType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFlowStatisticRequestLogType) UnmarshalJSON(b []byte) error {
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

type ListFlowStatisticRequestDirection struct {
	value string
}

type ListFlowStatisticRequestDirectionEnum struct {
	IN2OUT ListFlowStatisticRequestDirection
	OUT2IN ListFlowStatisticRequestDirection
}

func GetListFlowStatisticRequestDirectionEnum() ListFlowStatisticRequestDirectionEnum {
	return ListFlowStatisticRequestDirectionEnum{
		IN2OUT: ListFlowStatisticRequestDirection{
			value: "in2out",
		},
		OUT2IN: ListFlowStatisticRequestDirection{
			value: "out2in",
		},
	}
}

func (c ListFlowStatisticRequestDirection) Value() string {
	return c.value
}

func (c ListFlowStatisticRequestDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFlowStatisticRequestDirection) UnmarshalJSON(b []byte) error {
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

type ListFlowStatisticRequestAssetType struct {
	value string
}

type ListFlowStatisticRequestAssetTypeEnum struct {
	PUBLIC  ListFlowStatisticRequestAssetType
	PRIVATE ListFlowStatisticRequestAssetType
}

func GetListFlowStatisticRequestAssetTypeEnum() ListFlowStatisticRequestAssetTypeEnum {
	return ListFlowStatisticRequestAssetTypeEnum{
		PUBLIC: ListFlowStatisticRequestAssetType{
			value: "public",
		},
		PRIVATE: ListFlowStatisticRequestAssetType{
			value: "private",
		},
	}
}

func (c ListFlowStatisticRequestAssetType) Value() string {
	return c.value
}

func (c ListFlowStatisticRequestAssetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFlowStatisticRequestAssetType) UnmarshalJSON(b []byte) error {
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

type ListFlowStatisticRequestItem struct {
	value string
}

type ListFlowStatisticRequestItemEnum struct {
	SRC_IP          ListFlowStatisticRequestItem
	DST_IP          ListFlowStatisticRequestItem
	DST_PORT        ListFlowStatisticRequestItem
	PROTOCOL        ListFlowStatisticRequestItem
	DST_HOST        ListFlowStatisticRequestItem
	APP             ListFlowStatisticRequestItem
	RISK_IP         ListFlowStatisticRequestItem
	RISK_HOST       ListFlowStatisticRequestItem
	OPEN_PORT       ListFlowStatisticRequestItem
	DST_REGION_NAME ListFlowStatisticRequestItem
	SRC_REGION_NAME ListFlowStatisticRequestItem
}

func GetListFlowStatisticRequestItemEnum() ListFlowStatisticRequestItemEnum {
	return ListFlowStatisticRequestItemEnum{
		SRC_IP: ListFlowStatisticRequestItem{
			value: "src_ip",
		},
		DST_IP: ListFlowStatisticRequestItem{
			value: "dst_ip",
		},
		DST_PORT: ListFlowStatisticRequestItem{
			value: "dst_port",
		},
		PROTOCOL: ListFlowStatisticRequestItem{
			value: "protocol",
		},
		DST_HOST: ListFlowStatisticRequestItem{
			value: "dst_host",
		},
		APP: ListFlowStatisticRequestItem{
			value: "app",
		},
		RISK_IP: ListFlowStatisticRequestItem{
			value: "risk_ip",
		},
		RISK_HOST: ListFlowStatisticRequestItem{
			value: "risk_host",
		},
		OPEN_PORT: ListFlowStatisticRequestItem{
			value: "open_port",
		},
		DST_REGION_NAME: ListFlowStatisticRequestItem{
			value: "dst_region_name",
		},
		SRC_REGION_NAME: ListFlowStatisticRequestItem{
			value: "src_region_name",
		},
	}
}

func (c ListFlowStatisticRequestItem) Value() string {
	return c.value
}

func (c ListFlowStatisticRequestItem) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFlowStatisticRequestItem) UnmarshalJSON(b []byte) error {
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
