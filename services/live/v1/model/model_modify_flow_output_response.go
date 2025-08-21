package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ModifyFlowOutputResponse Response Object
type ModifyFlowOutputResponse struct {

	// ip白名单，最大20条ip白名单
	CidrWhitelist *[]string `json:"cidr_whitelist,omitempty"`

	// 推流地址，支持ip和域名，最大值64
	Destination *string `json:"destination,omitempty"`

	// output名称，有效字符：大小写字母，数字，中划线，下划线;且只能以字母开头
	Name *string `json:"name,omitempty"`

	Encryption *FlowSourceDecryption `json:"encryption,omitempty"`

	// 输出状态，ENABLED：启用，DISABLED：禁用
	OutputStatus *ModifyFlowOutputResponseOutputStatus `json:"output_status,omitempty"`

	// srt-listener模式，监听地址，不支持修改
	ListenerAddress *string `json:"listener_address,omitempty"`

	// 端口，如果ssrt-listener为监听端口，那么srt-caller为对端端口
	Port *int32 `json:"port,omitempty"`

	// 协议，srt-caller，srt-listener
	Protocol *ModifyFlowOutputResponseProtocol `json:"protocol,omitempty"`

	// srt-caller模式支持设置streamid
	StreamId *string `json:"stream_id,omitempty"`

	// 输出描述
	Description *string `json:"description,omitempty"`

	// 最小时延，单位毫秒，默认值2000，最小值10，最大值15000
	MinLatency *int32 `json:"min_latency,omitempty"`

	// 转推流状态，CONNECTED：转推中，DISCONNECTED：转推中断
	HealthStatus *ModifyFlowOutputResponseHealthStatus `json:"health_status,omitempty"`

	// 输出类型
	Type           *string `json:"type,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ModifyFlowOutputResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyFlowOutputResponse struct{}"
	}

	return strings.Join([]string{"ModifyFlowOutputResponse", string(data)}, " ")
}

type ModifyFlowOutputResponseOutputStatus struct {
	value string
}

type ModifyFlowOutputResponseOutputStatusEnum struct {
	ENABLED  ModifyFlowOutputResponseOutputStatus
	DISABLED ModifyFlowOutputResponseOutputStatus
}

func GetModifyFlowOutputResponseOutputStatusEnum() ModifyFlowOutputResponseOutputStatusEnum {
	return ModifyFlowOutputResponseOutputStatusEnum{
		ENABLED: ModifyFlowOutputResponseOutputStatus{
			value: "ENABLED",
		},
		DISABLED: ModifyFlowOutputResponseOutputStatus{
			value: "DISABLED",
		},
	}
}

func (c ModifyFlowOutputResponseOutputStatus) Value() string {
	return c.value
}

func (c ModifyFlowOutputResponseOutputStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyFlowOutputResponseOutputStatus) UnmarshalJSON(b []byte) error {
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

type ModifyFlowOutputResponseProtocol struct {
	value string
}

type ModifyFlowOutputResponseProtocolEnum struct {
	SRT_LISTENER ModifyFlowOutputResponseProtocol
	SRT_CALLER   ModifyFlowOutputResponseProtocol
}

func GetModifyFlowOutputResponseProtocolEnum() ModifyFlowOutputResponseProtocolEnum {
	return ModifyFlowOutputResponseProtocolEnum{
		SRT_LISTENER: ModifyFlowOutputResponseProtocol{
			value: "srt-listener",
		},
		SRT_CALLER: ModifyFlowOutputResponseProtocol{
			value: "srt-caller",
		},
	}
}

func (c ModifyFlowOutputResponseProtocol) Value() string {
	return c.value
}

func (c ModifyFlowOutputResponseProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyFlowOutputResponseProtocol) UnmarshalJSON(b []byte) error {
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

type ModifyFlowOutputResponseHealthStatus struct {
	value string
}

type ModifyFlowOutputResponseHealthStatusEnum struct {
	DISCONNECTED ModifyFlowOutputResponseHealthStatus
	CONNECTED    ModifyFlowOutputResponseHealthStatus
}

func GetModifyFlowOutputResponseHealthStatusEnum() ModifyFlowOutputResponseHealthStatusEnum {
	return ModifyFlowOutputResponseHealthStatusEnum{
		DISCONNECTED: ModifyFlowOutputResponseHealthStatus{
			value: "DISCONNECTED",
		},
		CONNECTED: ModifyFlowOutputResponseHealthStatus{
			value: "CONNECTED",
		},
	}
}

func (c ModifyFlowOutputResponseHealthStatus) Value() string {
	return c.value
}

func (c ModifyFlowOutputResponseHealthStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyFlowOutputResponseHealthStatus) UnmarshalJSON(b []byte) error {
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
