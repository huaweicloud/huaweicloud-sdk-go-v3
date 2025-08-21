package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type FlowsOutput struct {

	// ip白名单，最大20条ip白名单
	CidrWhitelist *[]string `json:"cidr_whitelist,omitempty"`

	// 推流地址，支持ip和域名，最大值64
	Destination *string `json:"destination,omitempty"`

	// output名称，有效字符：大小写字母，数字，中划线，下划线;且只能以字母开头
	Name string `json:"name"`

	Encryption *FlowSourceDecryption `json:"encryption,omitempty"`

	// 输出状态，ENABLED：启用，DISABLED：禁用
	OutputStatus *FlowsOutputOutputStatus `json:"output_status,omitempty"`

	// srt-listener模式，监听地址，不支持修改
	ListenerAddress *string `json:"listener_address,omitempty"`

	// 端口，如果ssrt-listener为监听端口，那么srt-caller为对端端口
	Port *int32 `json:"port,omitempty"`

	// 协议，srt-caller，srt-listener
	Protocol *FlowsOutputProtocol `json:"protocol,omitempty"`

	// srt-caller模式支持设置streamid
	StreamId *string `json:"stream_id,omitempty"`

	// 输出描述
	Description *string `json:"description,omitempty"`

	// 最小时延，单位毫秒，默认值2000，最小值10，最大值15000
	MinLatency *int32 `json:"min_latency,omitempty"`

	// 转推流状态，CONNECTED：转推中，DISCONNECTED：转推中断
	HealthStatus *FlowsOutputHealthStatus `json:"health_status,omitempty"`

	// 输出类型
	Type *string `json:"type,omitempty"`
}

func (o FlowsOutput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlowsOutput struct{}"
	}

	return strings.Join([]string{"FlowsOutput", string(data)}, " ")
}

type FlowsOutputOutputStatus struct {
	value string
}

type FlowsOutputOutputStatusEnum struct {
	ENABLED  FlowsOutputOutputStatus
	DISABLED FlowsOutputOutputStatus
}

func GetFlowsOutputOutputStatusEnum() FlowsOutputOutputStatusEnum {
	return FlowsOutputOutputStatusEnum{
		ENABLED: FlowsOutputOutputStatus{
			value: "ENABLED",
		},
		DISABLED: FlowsOutputOutputStatus{
			value: "DISABLED",
		},
	}
}

func (c FlowsOutputOutputStatus) Value() string {
	return c.value
}

func (c FlowsOutputOutputStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlowsOutputOutputStatus) UnmarshalJSON(b []byte) error {
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

type FlowsOutputProtocol struct {
	value string
}

type FlowsOutputProtocolEnum struct {
	SRT_LISTENER FlowsOutputProtocol
	SRT_CALLER   FlowsOutputProtocol
}

func GetFlowsOutputProtocolEnum() FlowsOutputProtocolEnum {
	return FlowsOutputProtocolEnum{
		SRT_LISTENER: FlowsOutputProtocol{
			value: "srt-listener",
		},
		SRT_CALLER: FlowsOutputProtocol{
			value: "srt-caller",
		},
	}
}

func (c FlowsOutputProtocol) Value() string {
	return c.value
}

func (c FlowsOutputProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlowsOutputProtocol) UnmarshalJSON(b []byte) error {
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

type FlowsOutputHealthStatus struct {
	value string
}

type FlowsOutputHealthStatusEnum struct {
	DISCONNECTED FlowsOutputHealthStatus
	CONNECTED    FlowsOutputHealthStatus
}

func GetFlowsOutputHealthStatusEnum() FlowsOutputHealthStatusEnum {
	return FlowsOutputHealthStatusEnum{
		DISCONNECTED: FlowsOutputHealthStatus{
			value: "DISCONNECTED",
		},
		CONNECTED: FlowsOutputHealthStatus{
			value: "CONNECTED",
		},
	}
}

func (c FlowsOutputHealthStatus) Value() string {
	return c.value
}

func (c FlowsOutputHealthStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlowsOutputHealthStatus) UnmarshalJSON(b []byte) error {
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
