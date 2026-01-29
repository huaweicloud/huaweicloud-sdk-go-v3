package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FlowSource 入流信息
type FlowSource struct {

	// 拉流地址
	SourceListenerAddress *string `json:"source_listener_address,omitempty"`

	// 拉流端口，2077/2088不可选
	SourceListenerPort *int32 `json:"source_listener_port,omitempty"`

	// srt拉流streamid
	StreamId *string `json:"stream_id,omitempty"`

	// 最小时延，单位ms
	MinLatency *int32 `json:"min_latency,omitempty"`

	// 推流CIDR IP白名单列表
	CidrWhitelist *[]string `json:"cidr_whitelist,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 协议，srt-caller，srt-listener
	Protocol FlowSourceProtocol `json:"protocol"`

	// 入流资源名称
	Name string `json:"name"`

	Decryption *FlowSourceDecryption `json:"decryption,omitempty"`

	// **参数解释**： 转推流状态 **约束限制**： 不涉及 **取值范围**： - CONNECTED：转推中 - DISCONNECTED：转推中断
	HealthStatus *FlowSourceHealthStatus `json:"health_status,omitempty"`
}

func (o FlowSource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlowSource struct{}"
	}

	return strings.Join([]string{"FlowSource", string(data)}, " ")
}

type FlowSourceProtocol struct {
	value string
}

type FlowSourceProtocolEnum struct {
	SRT_CALLER   FlowSourceProtocol
	SRT_LISTENER FlowSourceProtocol
}

func GetFlowSourceProtocolEnum() FlowSourceProtocolEnum {
	return FlowSourceProtocolEnum{
		SRT_CALLER: FlowSourceProtocol{
			value: "srt-caller",
		},
		SRT_LISTENER: FlowSourceProtocol{
			value: "srt-listener",
		},
	}
}

func (c FlowSourceProtocol) Value() string {
	return c.value
}

func (c FlowSourceProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlowSourceProtocol) UnmarshalJSON(b []byte) error {
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

type FlowSourceHealthStatus struct {
	value string
}

type FlowSourceHealthStatusEnum struct {
	CONNECTED    FlowSourceHealthStatus
	DISCONNECTED FlowSourceHealthStatus
}

func GetFlowSourceHealthStatusEnum() FlowSourceHealthStatusEnum {
	return FlowSourceHealthStatusEnum{
		CONNECTED: FlowSourceHealthStatus{
			value: "CONNECTED",
		},
		DISCONNECTED: FlowSourceHealthStatus{
			value: "DISCONNECTED",
		},
	}
}

func (c FlowSourceHealthStatus) Value() string {
	return c.value
}

func (c FlowSourceHealthStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlowSourceHealthStatus) UnmarshalJSON(b []byte) error {
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
