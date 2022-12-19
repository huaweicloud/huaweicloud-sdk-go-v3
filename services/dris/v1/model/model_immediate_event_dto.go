package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// **参数说明**：即时事件的内容。
type ImmediateEventDto struct {

	// **参数说明**：事件发生时间，毫秒级。  格式：yyyy-MM-dd''T''HH:mm:ss.SSS''Z''  例如 2015-12-12T12:12:12.356Z。
	TimeStamp *sdktime.SdkTime `json:"time_stamp"`

	// **参数说明**：事件类型，参照附录[《国标交通事件及标志列表》](https://support.huaweicloud.com/api-v2x/v2x_04_0101.html)文件。当填写event_type时，event_class为必选。
	EventClass ImmediateEventDtoEventClass `json:"event_class"`

	// **参数说明**：事件类型，参照附录[《国标交通事件及标志列表》](https://support.huaweicloud.com/api-v2x/v2x_04_0101.html)文件。
	EventType int32 `json:"event_type"`

	// **参数说明**：事件来源类型列表,支持事件来源。  **取值范围**：  - unknown：未知数据  - police：警方数据  - government：政府数据  - meteorological：气象数据  - internet：互联网数据  - detection：检测器检测到的数据  - v2xServer：平台上报数据  - rsu：RSU上报数据  - obu：车载终端上报数据
	EventSource ImmediateEventDtoEventSource `json:"event_source"`

	// **参数说明**：事件来源的ID，由用户自定义。
	EventSourceId *string `json:"event_source_id,omitempty"`

	// **参数说明**：道路交通事件的信息来源提供的事件置信度水平，帮助接收端判断是否相信该事件信息，可选。
	EventConfidence *int32 `json:"event_confidence,omitempty"`

	EventPosition *ImmediateEventPosition3D `json:"event_position"`

	// **参数说明**：事件的发生区域半径，单位为分米。
	EventRadius *int32 `json:"event_radius,omitempty"`

	// **参数说明**：事件的文本描述信息,可自行扩展需传递的信息。
	EventDescription *string `json:"event_description,omitempty"`

	// **参数说明**：事件优先级，0-7，数字越大，级别越高。
	EventPriority *int32 `json:"event_priority,omitempty"`

	// **参数说明**：事件生效的关联路径。
	ReferencePaths *[]ImmediateEventReferencePath `json:"reference_paths,omitempty"`
}

func (o ImmediateEventDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImmediateEventDto struct{}"
	}

	return strings.Join([]string{"ImmediateEventDto", string(data)}, " ")
}

type ImmediateEventDtoEventClass struct {
	value string
}

type ImmediateEventDtoEventClassEnum struct {
	ABNORMAL_TRAFFIC ImmediateEventDtoEventClass
	ADVERSE_WEATHER  ImmediateEventDtoEventClass
	ABNORMAL_VEHICLE ImmediateEventDtoEventClass
	TRAFFIC_SIGN     ImmediateEventDtoEventClass
}

func GetImmediateEventDtoEventClassEnum() ImmediateEventDtoEventClassEnum {
	return ImmediateEventDtoEventClassEnum{
		ABNORMAL_TRAFFIC: ImmediateEventDtoEventClass{
			value: "abnormal traffic",
		},
		ADVERSE_WEATHER: ImmediateEventDtoEventClass{
			value: "adverse weather",
		},
		ABNORMAL_VEHICLE: ImmediateEventDtoEventClass{
			value: "abnormal vehicle",
		},
		TRAFFIC_SIGN: ImmediateEventDtoEventClass{
			value: "traffic sign",
		},
	}
}

func (c ImmediateEventDtoEventClass) Value() string {
	return c.value
}

func (c ImmediateEventDtoEventClass) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImmediateEventDtoEventClass) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ImmediateEventDtoEventSource struct {
	value string
}

type ImmediateEventDtoEventSourceEnum struct {
	UNKNOWN        ImmediateEventDtoEventSource
	POLICE         ImmediateEventDtoEventSource
	GOVERNMENT     ImmediateEventDtoEventSource
	METEOROLOGICAL ImmediateEventDtoEventSource
	INTERNET       ImmediateEventDtoEventSource
	DETECTION      ImmediateEventDtoEventSource
	V2X_SERVER     ImmediateEventDtoEventSource
	RSU            ImmediateEventDtoEventSource
	OBU            ImmediateEventDtoEventSource
}

func GetImmediateEventDtoEventSourceEnum() ImmediateEventDtoEventSourceEnum {
	return ImmediateEventDtoEventSourceEnum{
		UNKNOWN: ImmediateEventDtoEventSource{
			value: "unknown",
		},
		POLICE: ImmediateEventDtoEventSource{
			value: "police",
		},
		GOVERNMENT: ImmediateEventDtoEventSource{
			value: "government",
		},
		METEOROLOGICAL: ImmediateEventDtoEventSource{
			value: "meteorological",
		},
		INTERNET: ImmediateEventDtoEventSource{
			value: "internet",
		},
		DETECTION: ImmediateEventDtoEventSource{
			value: "detection",
		},
		V2X_SERVER: ImmediateEventDtoEventSource{
			value: "v2xServer",
		},
		RSU: ImmediateEventDtoEventSource{
			value: "rsu",
		},
		OBU: ImmediateEventDtoEventSource{
			value: "obu",
		},
	}
}

func (c ImmediateEventDtoEventSource) Value() string {
	return c.value
}

func (c ImmediateEventDtoEventSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImmediateEventDtoEventSource) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
