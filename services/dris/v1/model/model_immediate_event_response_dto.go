package model

import (
	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ImmediateEventResponseDto **参数说明**：即时事件的内容。
type ImmediateEventResponseDto struct {

	// **参数说明**：事件发生时间，毫秒级。  格式：yyyy-MM-dd''T''HH:mm:ss.SSS''Z''  例如 2015-12-12T12:12:12.356Z。
	TimeStamp *sdktime.SdkTime `json:"time_stamp,omitempty"`

	// **参数说明**：事件类型，参照附录[《国标交通事件及标志列表》](https://support.huaweicloud.com/api-v2x/v2x_04_0101.html)文件。当填写event_type时，event_class为必选。
	EventClass *ImmediateEventResponseDtoEventClass `json:"event_class,omitempty"`

	// **参数说明**：事件类型，参照附录[《国标交通事件及标志列表》](https://support.huaweicloud.com/api-v2x/v2x_04_0101.html)文件
	EventType *int32 `json:"event_type,omitempty"`

	// **参数说明**：事件来源类型列表,支持事件来源。  **取值范围**：  - unknown：未知数据  - police：警方数据  - government：政府数据  - meteorological：气象数据  - internet：互联网数据  - detection：检测器检测到的数据  - v2xServer：平台上报数据  - rsu：RSU上报数据  - obu：车载终端上报数据
	EventSource *ImmediateEventResponseDtoEventSource `json:"event_source,omitempty"`

	// **参数说明**：事件来源的ID，由用户自定义。
	EventSourceId *string `json:"event_source_id,omitempty"`

	// **参数说明**：道路交通事件的信息来源提供的事件置信度水平，帮助接收端判断是否相信该事件信息，可选。
	EventConfidence *int32 `json:"event_confidence,omitempty"`

	EventPosition *ImmediateEventPosition3D `json:"event_position,omitempty"`

	// **参数说明**：事件的发生区域半径，单位为分米。
	EventRadius *int32 `json:"event_radius,omitempty"`

	// **参数说明**：事件的文本描述信息,可自行扩展需传递的信息。
	EventDescription *string `json:"event_description,omitempty"`

	// **参数说明**：事件优先级，0-7，数字越大，级别越高。
	EventPriority *int32 `json:"event_priority,omitempty"`

	// **参数说明**：事件生效的关联路径。
	ReferencePaths *[]ImmediateEventReferencePath `json:"reference_paths,omitempty"`
}

func (o ImmediateEventResponseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImmediateEventResponseDto struct{}"
	}

	return strings.Join([]string{"ImmediateEventResponseDto", string(data)}, " ")
}

type ImmediateEventResponseDtoEventClass struct {
	value string
}

type ImmediateEventResponseDtoEventClassEnum struct {
	ABNORMAL_TRAFFIC ImmediateEventResponseDtoEventClass
	ADVERSE_WEATHER  ImmediateEventResponseDtoEventClass
	ABNORMAL_VEHICLE ImmediateEventResponseDtoEventClass
	TRAFFIC_SIGN     ImmediateEventResponseDtoEventClass
}

func GetImmediateEventResponseDtoEventClassEnum() ImmediateEventResponseDtoEventClassEnum {
	return ImmediateEventResponseDtoEventClassEnum{
		ABNORMAL_TRAFFIC: ImmediateEventResponseDtoEventClass{
			value: "abnormal traffic",
		},
		ADVERSE_WEATHER: ImmediateEventResponseDtoEventClass{
			value: "adverse weather",
		},
		ABNORMAL_VEHICLE: ImmediateEventResponseDtoEventClass{
			value: "abnormal vehicle",
		},
		TRAFFIC_SIGN: ImmediateEventResponseDtoEventClass{
			value: "traffic sign",
		},
	}
}

func (c ImmediateEventResponseDtoEventClass) Value() string {
	return c.value
}

func (c ImmediateEventResponseDtoEventClass) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImmediateEventResponseDtoEventClass) UnmarshalJSON(b []byte) error {
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

type ImmediateEventResponseDtoEventSource struct {
	value string
}

type ImmediateEventResponseDtoEventSourceEnum struct {
	UNKNOWN        ImmediateEventResponseDtoEventSource
	POLICE         ImmediateEventResponseDtoEventSource
	GOVERNMENT     ImmediateEventResponseDtoEventSource
	METEOROLOGICAL ImmediateEventResponseDtoEventSource
	INTERNET       ImmediateEventResponseDtoEventSource
	DETECTION      ImmediateEventResponseDtoEventSource
	V2X_SERVER     ImmediateEventResponseDtoEventSource
	RSU            ImmediateEventResponseDtoEventSource
	OBU            ImmediateEventResponseDtoEventSource
}

func GetImmediateEventResponseDtoEventSourceEnum() ImmediateEventResponseDtoEventSourceEnum {
	return ImmediateEventResponseDtoEventSourceEnum{
		UNKNOWN: ImmediateEventResponseDtoEventSource{
			value: "unknown",
		},
		POLICE: ImmediateEventResponseDtoEventSource{
			value: "police",
		},
		GOVERNMENT: ImmediateEventResponseDtoEventSource{
			value: "government",
		},
		METEOROLOGICAL: ImmediateEventResponseDtoEventSource{
			value: "meteorological",
		},
		INTERNET: ImmediateEventResponseDtoEventSource{
			value: "internet",
		},
		DETECTION: ImmediateEventResponseDtoEventSource{
			value: "detection",
		},
		V2X_SERVER: ImmediateEventResponseDtoEventSource{
			value: "v2xServer",
		},
		RSU: ImmediateEventResponseDtoEventSource{
			value: "rsu",
		},
		OBU: ImmediateEventResponseDtoEventSource{
			value: "obu",
		},
	}
}

func (c ImmediateEventResponseDtoEventSource) Value() string {
	return c.value
}

func (c ImmediateEventResponseDtoEventSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImmediateEventResponseDtoEventSource) UnmarshalJSON(b []byte) error {
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
