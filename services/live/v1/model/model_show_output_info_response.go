package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowOutputInfoResponse Response Object
type ShowOutputInfoResponse struct {

	// ip白名单，最大20条ip白名单
	CidrWhitelist *[]string `json:"cidr_whitelist,omitempty"`

	// 推流地址，支持ip和域名，最大值64
	Destination *string `json:"destination,omitempty"`

	// output名称，有效字符：大小写字母，数字，中划线，下划线;且只能以字母开头
	Name *string `json:"name,omitempty"`

	Encryption *FlowSourceDecryption `json:"encryption,omitempty"`

	// 输出状态，ENABLED：启用，DISABLED：禁用
	OutputStatus *ShowOutputInfoResponseOutputStatus `json:"output_status,omitempty"`

	// srt-listener模式，监听地址，不支持修改
	ListenerAddress *string `json:"listener_address,omitempty"`

	// 端口，如果ssrt-listener为监听端口，那么srt-caller为对端端口
	Port *int32 `json:"port,omitempty"`

	// 协议，srt-caller，srt-listener
	Protocol *ShowOutputInfoResponseProtocol `json:"protocol,omitempty"`

	// srt-caller模式支持设置streamid
	StreamId *string `json:"stream_id,omitempty"`

	// 输出描述
	Description *string `json:"description,omitempty"`

	// 最小时延，单位毫秒，默认值2000，最小值10，最大值15000
	MinLatency *int32 `json:"min_latency,omitempty"`

	// 转推流状态，CONNECTED：转推中，DISCONNECTED：转推中断
	HealthStatus *ShowOutputInfoResponseHealthStatus `json:"health_status,omitempty"`

	// 输出类型
	Type           *string `json:"type,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowOutputInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOutputInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowOutputInfoResponse", string(data)}, " ")
}

type ShowOutputInfoResponseOutputStatus struct {
	value string
}

type ShowOutputInfoResponseOutputStatusEnum struct {
	ENABLED  ShowOutputInfoResponseOutputStatus
	DISABLED ShowOutputInfoResponseOutputStatus
}

func GetShowOutputInfoResponseOutputStatusEnum() ShowOutputInfoResponseOutputStatusEnum {
	return ShowOutputInfoResponseOutputStatusEnum{
		ENABLED: ShowOutputInfoResponseOutputStatus{
			value: "ENABLED",
		},
		DISABLED: ShowOutputInfoResponseOutputStatus{
			value: "DISABLED",
		},
	}
}

func (c ShowOutputInfoResponseOutputStatus) Value() string {
	return c.value
}

func (c ShowOutputInfoResponseOutputStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowOutputInfoResponseOutputStatus) UnmarshalJSON(b []byte) error {
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

type ShowOutputInfoResponseProtocol struct {
	value string
}

type ShowOutputInfoResponseProtocolEnum struct {
	SRT_LISTENER ShowOutputInfoResponseProtocol
	SRT_CALLER   ShowOutputInfoResponseProtocol
}

func GetShowOutputInfoResponseProtocolEnum() ShowOutputInfoResponseProtocolEnum {
	return ShowOutputInfoResponseProtocolEnum{
		SRT_LISTENER: ShowOutputInfoResponseProtocol{
			value: "srt-listener",
		},
		SRT_CALLER: ShowOutputInfoResponseProtocol{
			value: "srt-caller",
		},
	}
}

func (c ShowOutputInfoResponseProtocol) Value() string {
	return c.value
}

func (c ShowOutputInfoResponseProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowOutputInfoResponseProtocol) UnmarshalJSON(b []byte) error {
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

type ShowOutputInfoResponseHealthStatus struct {
	value string
}

type ShowOutputInfoResponseHealthStatusEnum struct {
	DISCONNECTED ShowOutputInfoResponseHealthStatus
	CONNECTED    ShowOutputInfoResponseHealthStatus
}

func GetShowOutputInfoResponseHealthStatusEnum() ShowOutputInfoResponseHealthStatusEnum {
	return ShowOutputInfoResponseHealthStatusEnum{
		DISCONNECTED: ShowOutputInfoResponseHealthStatus{
			value: "DISCONNECTED",
		},
		CONNECTED: ShowOutputInfoResponseHealthStatus{
			value: "CONNECTED",
		},
	}
}

func (c ShowOutputInfoResponseHealthStatus) Value() string {
	return c.value
}

func (c ShowOutputInfoResponseHealthStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowOutputInfoResponseHealthStatus) UnmarshalJSON(b []byte) error {
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
